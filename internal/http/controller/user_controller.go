package controller

import (
	"context"
	"encoding/json"
	"github.com/arvians-id/go-gorm/internal/http/presenter/request"
	"github.com/arvians-id/go-gorm/internal/http/presenter/response"
	"github.com/arvians-id/go-gorm/internal/model"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		DB: db,
	}
}

func (controller *UserController) Route() http.Handler {
	route := chi.NewRouter()
	route.Get("/", controller.List)
	route.Get("/{id}", controller.FindById)
	route.Post("/", controller.Create)
	route.Patch("/{id}", controller.Update)
	route.Delete("/{id}", controller.Delete)
	route.Patch("/change-roles", controller.ChangeRoles)

	return route
}

func (controller *UserController) List(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var users []model.User
	err := controller.DB.WithContext(ctx).
		Table("user_roles").
		Select("COUNT(*) as total,*").
		Joins("LEFT JOIN roles r on r.id = user_roles.role_id").
		Joins("LEFT JOIN users u on u.id = user_roles.user_id").
		Group("u.id").
		Order("total desc").
		Preload("Roles").
		Find(&users).Error
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", users)
}

func (controller *UserController) FindById(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	id := chi.URLParam(r, "id")
	idUser, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	var userResponse model.User
	err = controller.DB.WithContext(ctx).Model(&model.User{
		ID: uint(idUser),
	}).Preload("Roles").First(&userResponse).Error
	if err != nil {
		response.ReturnErrorNotFound(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", userResponse)
}

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	err = controller.DB.Transaction(func(tx *gorm.DB) error {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		err = tx.WithContext(ctx).Create(&user).Association("Roles").Append(&model.Role{ID: 2, Role: "member"})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", user)
}

func (controller *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	id := chi.URLParam(r, "id")
	var userDB model.User
	err = controller.DB.WithContext(ctx).Preload("Roles").First(&userDB, id).Error
	if err != nil {
		response.ReturnErrorNotFound(w, err, nil)
		return
	}

	err = controller.DB.WithContext(ctx).Model(&userDB).Updates(&user).Error
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", userDB)
}

func (controller *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	id := chi.URLParam(r, "id")
	var user model.User
	err := controller.DB.WithContext(ctx).First(&user, id).Error
	if err != nil {
		response.ReturnErrorNotFound(w, err, nil)
		return
	}

	err = controller.DB.WithContext(ctx).Delete(&user).Error
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", user)
}

func (controller *UserController) ChangeRoles(w http.ResponseWriter, r *http.Request) {
	var userRequest request.ChangeRolesRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var roles []model.Role
	for _, role := range userRequest.RoleID {
		roles = append(roles, model.Role{ID: role})
	}

	err = controller.DB.WithContext(ctx).Model(&model.User{
		ID: userRequest.UserID,
	}).Association("Roles").Replace(roles)
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", userRequest)
}
