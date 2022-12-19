package service

import (
	"context"

	"github.com/arvians-id/go-gorm/internal/model"
	"github.com/arvians-id/go-gorm/internal/repository"
)

type UserService interface {
	List(ctx context.Context) ([]*model.User, error)
	FindById(ctx context.Context, id uint64) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint64) error
	UpdateRoles(ctx context.Context, userId uint64, roles []*model.Role) error
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *UserServiceImpl) List(ctx context.Context) ([]*model.User, error) {
	return service.UserRepository.List(ctx)
}

func (service *UserServiceImpl) FindById(ctx context.Context, id uint64) (*model.User, error) {
	return service.UserRepository.FindById(ctx, id)
}

func (service *UserServiceImpl) Create(ctx context.Context, user *model.User) (*model.User, error) {
	return service.UserRepository.Create(ctx, user)
}

func (service *UserServiceImpl) Update(ctx context.Context, user *model.User) error {
	return service.UserRepository.Update(ctx, user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, id uint64) error {
	return service.UserRepository.Delete(ctx, id)
}

func (service *UserServiceImpl) UpdateRoles(ctx context.Context, userId uint64, roles []*model.Role) error {
	return service.UserRepository.UpdateRoles(ctx, userId, roles)
}
