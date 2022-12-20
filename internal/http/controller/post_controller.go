package controller

import (
	"context"
	"encoding/json"
	"github.com/arvians-id/go-gorm/internal/http/presenter/request"
	"github.com/arvians-id/go-gorm/internal/http/presenter/response"
	"github.com/arvians-id/go-gorm/internal/model"
	"github.com/arvians-id/go-gorm/internal/service"
	"github.com/go-chi/chi/v5"
	"math"
	"net/http"
	"strconv"
	"time"
)

type PostController struct {
	PostService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{
		PostService: postService,
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
	perPage := 3
	pages := 1
	var offset int

	queryPage := r.URL.Query().Get("page")
	if queryPage != "" {
		page, err := strconv.ParseInt(queryPage, 10, 64)
		if err != nil {
			response.ReturnErrorBadRequest(w, err, nil)
			return
		}
		pages = int(page)
	}

	totalRows, err := controller.PostService.TotalRows(ctx)
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

	posts, err := controller.PostService.List(ctx, perPage, offset)
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
	idPost, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	post, err := controller.PostService.FindById(ctx, uint64(idPost))
	if err != nil {
		response.ReturnErrorNotFound(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", post)
}

func (controller *PostController) Create(w http.ResponseWriter, r *http.Request) {
	var postRequest *model.Post
	err := json.NewDecoder(r.Body).Decode(&postRequest)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	post, err := controller.PostService.Create(ctx, postRequest)
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
	idPost, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.ReturnErrorBadRequest(w, err, nil)
		return
	}

	err = controller.PostService.Delete(ctx, idPost)
	if err != nil {
		response.ReturnErrorInternalServerError(w, err, nil)
		return
	}

	response.ReturnSuccessOK(w, "success", nil)
}
