package server

import (
	"github.com/arvians-id/go-gorm/cmd/config"
	"github.com/arvians-id/go-gorm/internal/http/controller"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func NewInitializedDatabaseGorm(configuration config.Config) (*gorm.DB, error) {
	db, err := config.NewSQLiteGorm(configuration)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewInitializedServer(configuration config.Config) *chi.Mux {
	// Setup Configuration
	r := chi.NewRouter()

	// Setup Database
	db, err := NewInitializedDatabaseGorm(configuration)
	if err != nil {
		log.Fatal(err)
	}

	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	userController := controller.NewUserController(db)
	postController := controller.NewPostController(db)
	commentController := controller.NewCommentController(db)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/users", userController.Route())
		r.Mount("/posts", postController.Route())
		r.Mount("/comments", commentController.Route())
	})

	return r
}
