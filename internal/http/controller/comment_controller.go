package controller

import (
	"context"
	"encoding/json"
	"github.com/arvians-id/go-gorm/internal/http/presenter/response"
	"github.com/arvians-id/go-gorm/internal/model"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type CommentController struct {
	DB *gorm.DB
}

func NewCommentController(db *gorm.DB) *CommentController {
	return &CommentController{
		DB: db,
	}
}

func (controller *CommentController) Route() http.Handler {
	route := chi.NewRouter()
	route.Get("/", controller.List)
	route.Post("/", controller.Create)
	route.Get("/{id}", controller.FindById)
	route.Delete("/{id}", controller.Delete)

	return route
}

func (controller *CommentController) List(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var comments []model.Comment
	err := controller.DB.WithContext(ctx).Find(&comments).Error
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", comments)
}

func (controller *CommentController) FindById(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	id := chi.URLParam(r, "id")
	var comment model.Comment
	err := controller.DB.WithContext(ctx).First(&comment, id).Error
	if err != nil {
		response.ReturnErrorNotFound(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", comment)
}

func (controller *CommentController) Create(w http.ResponseWriter, r *http.Request) {
	var comment model.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err = controller.DB.WithContext(ctx).Create(&comment).Error
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", comment)
}

func (controller *CommentController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	id := chi.URLParam(r, "id")
	var comment model.Comment
	err := controller.DB.WithContext(ctx).First(&comment, id).Error
	if err != nil {
		response.ReturnErrorNotFound(w, err, nil)
		return
	}

	err = controller.DB.WithContext(ctx).Delete(&comment).Error
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", nil)
}
