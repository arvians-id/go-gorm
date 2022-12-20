package service

import (
	"context"

	"github.com/arvians-id/go-gorm/internal/model"
	"github.com/arvians-id/go-gorm/internal/repository"
)

type UserServiceContract interface {
	List(ctx context.Context) ([]*model.User, error)
	FindById(ctx context.Context, id uint64) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint64) error
	UpdateRoles(ctx context.Context, userId uint64, roles []*model.Role) error
}

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{
		UserRepository: userRepository,
	}
}

func (service *UserService) List(ctx context.Context) ([]*model.User, error) {
	return service.UserRepository.List(ctx)
}

func (service *UserService) FindById(ctx context.Context, id uint64) (*model.User, error) {
	return service.UserRepository.FindById(ctx, id)
}

func (service *UserService) Create(ctx context.Context, user *model.User) (*model.User, error) {
	return service.UserRepository.Create(ctx, user)
}

func (service *UserService) Update(ctx context.Context, user *model.User) error {
	return service.UserRepository.Update(ctx, user)
}

func (service *UserService) Delete(ctx context.Context, id uint64) error {
	return service.UserRepository.Delete(ctx, id)
}

func (service *UserService) UpdateRoles(ctx context.Context, userId uint64, roles []*model.Role) error {
	return service.UserRepository.UpdateRoles(ctx, userId, roles)
}
