package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/arvians-id/go-gorm/internal/http/presenter/request"
	"github.com/arvians-id/go-gorm/internal/http/presenter/response"
	"github.com/arvians-id/go-gorm/internal/model"
	"github.com/arvians-id/go-gorm/internal/service"
	"github.com/go-chi/chi/v5"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (controller *UserController) Routes() http.Handler {
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

	users, err := controller.UserService.List(ctx)
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
	idUser, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	user, err := controller.UserService.FindById(ctx, idUser)

	response.ReturnSuccessOK(w, "success", user)
}

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var user *model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	userCreated, err := controller.UserService.Create(ctx, user)
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", userCreated)
}

func (controller *UserController) Update(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var user *model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	id := chi.URLParam(r, "id")
	idUser, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	user.ID = idUser
	err = controller.UserService.Update(ctx, user)
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", user)
}

func (controller *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	id := chi.URLParam(r, "id")
	idUser, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	err = controller.UserService.Delete(ctx, idUser)
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", nil)
}

func (controller *UserController) ChangeRoles(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var userRequest request.ChangeRolesRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	var roles []*model.Role
	for _, role := range userRequest.RoleID {
		roles = append(roles, &model.Role{ID: role})
	}

	err = controller.UserService.UpdateRoles(ctx, userRequest.UserID, roles)
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", nil)
}
