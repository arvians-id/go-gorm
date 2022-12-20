package controller

import (
	"context"
	"encoding/json"
	"github.com/arvians-id/go-gorm/internal/http/presenter/response"
	"github.com/arvians-id/go-gorm/internal/model"
	"github.com/arvians-id/go-gorm/internal/service"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"
)

type CommentController struct {
	CommentService service.CommentService
}

func NewCommentController(commentService service.CommentService) *CommentController {
	return &CommentController{
		CommentService: commentService,
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

	comments, err := controller.CommentService.List(ctx)
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
	idComment, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	comment, err := controller.CommentService.FindById(ctx, idComment)
	if err != nil {
		response.ReturnErrorNotFound(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", comment)
}

func (controller *CommentController) Create(w http.ResponseWriter, r *http.Request) {
	var commentRequest *model.Comment
	err := json.NewDecoder(r.Body).Decode(&commentRequest)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	comment, err := controller.CommentService.Create(ctx, commentRequest)
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
	idComment, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	err = controller.CommentService.Delete(ctx, idComment)
	if err != nil {
		response.ReturnErrorNotFound(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", nil)
}
