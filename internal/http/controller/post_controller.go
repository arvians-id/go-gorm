package controller

import (
	"context"
	"encoding/json"
	"github.com/arvians-id/go-gorm/internal/http/presenter/request"
	"github.com/arvians-id/go-gorm/internal/http/presenter/response"
	"github.com/arvians-id/go-gorm/internal/model"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"net/http"
	"strconv"
	"time"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(db *gorm.DB) *PostController {
	return &PostController{
		DB: db,
	}
}

func (controller *PostController) Route() http.Handler {
	route := chi.NewRouter()
	route.Get("/", controller.List)
	route.Get("/{id}", controller.FindById)
	route.Post("/", controller.Create)
	route.Delete("/{id}", controller.Delete)

	return route
}

func (controller *PostController) List(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var pagination request.PaginationData
	queryPage := r.URL.Query().Get("page")
	perPage := 3
	pages := 1
	var offset int
	var totalRows int64

	if queryPage != "" {
		page, err := strconv.ParseInt(queryPage, 10, 64)
		if err != nil {
			response.ReturnErrorBadRequest(w, err, nil)
			return
		}
		pages = int(page)
	}

	err := controller.DB.WithContext(ctx).Model(&model.Post{}).Count(&totalRows).Error
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	pagination.NextPage = pages + 1
	pagination.PreviousPage = pages - 1
	pagination.CurrentPage = pages
	pagination.TotalPage = int(math.Ceil(float64(totalRows) / float64(perPage)))
	pagination.TotalData = int(totalRows)

	offset = (pages - 1) * perPage

	var posts []model.Post
	err = controller.DB.WithContext(ctx).Preload(clause.Associations, func(db *gorm.DB) *gorm.DB {
		return db.Order("comments.created_at desc")
	}).Order("posts.created_at desc").Limit(perPage).Offset(offset).Find(&posts).Error
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	response.ReturnSuccessPagesOK(w, "success", posts, pagination)
}

func (controller *PostController) FindById(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	id := chi.URLParam(r, "id")
	var post model.Post
	err := controller.DB.WithContext(ctx).Preload(clause.Associations, func(db *gorm.DB) *gorm.DB {
		return db.Order("comments.created_at desc")
	}).First(&post, id).Error
	if err != nil {
		response.ReturnErrorNotFound(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", post)
}

func (controller *PostController) Create(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err = controller.DB.WithContext(ctx).Create(&post).Error
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", post)
}

func (controller *PostController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	id := chi.URLParam(r, "id")
	var post model.Post
	err := controller.DB.WithContext(ctx).First(&post, id).Error
	if err != nil {
		response.ReturnErrorNotFound(w, err, nil)
		return
	}

	err = controller.DB.WithContext(ctx).Delete(&post).Error
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", post)
}
