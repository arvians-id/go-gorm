package server

import (
	"github.com/arvians-id/go-gorm/internal/http/controller"
	"github.com/go-chi/chi/v5"
)

func NewRoutes(userController controller.UserController, postController controller.PostController, commentController controller.CommentController) *chi.Mux {
	// Setup Configuration
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Mount("/users", userController.Routes())
		r.Mount("/posts", postController.Routes())
		r.Mount("/comments", commentController.Routes())
	})

	return r
}
