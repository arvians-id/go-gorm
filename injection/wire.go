//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/arvians-id/go-gorm/cmd/config"
	"github.com/arvians-id/go-gorm/cmd/server"
	"github.com/arvians-id/go-gorm/internal/http/controller"
	"github.com/arvians-id/go-gorm/internal/repository"
	"github.com/arvians-id/go-gorm/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
	controller.NewUserController,
)

var postSet = wire.NewSet(
	repository.NewPostRepository,
	service.NewPostService,
	controller.NewPostController,
)

var commentSet = wire.NewSet(
	repository.NewCommentRepository,
	service.NewCommentService,
	controller.NewCommentController,
)

func InitServerAPI(configuration config.Config) (*chi.Mux, error) {
	wire.Build(
		config.NewSQLiteGorm,
		userSet,
		postSet,
		commentSet,
		server.NewRoutes,
	)

	return nil, nil
}
