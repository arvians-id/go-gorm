package service

import (
	"context"

	"github.com/arvians-id/go-gorm/internal/model"
	"github.com/arvians-id/go-gorm/internal/repository"
)

type PostServiceContract interface {
	List(ctx context.Context, perPage int, offset int) ([]*model.Post, error)
	FindById(ctx context.Context, id uint64) (*model.PostResponse, error)
	Create(ctx context.Context, post *model.Post) (*model.Post, error)
	Delete(ctx context.Context, id uint64) error
	TotalRows(ctx context.Context) (int64, error)
}

type PostService struct {
	PostRepository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) PostService {
	return PostService{
		PostRepository: postRepository,
	}
}

func (service *PostService) List(ctx context.Context, perPage int, offset int) ([]*model.Post, error) {
	return service.PostRepository.List(ctx, perPage, offset)
}

func (service *PostService) FindById(ctx context.Context, id uint64) (*model.PostResponse, error) {
	return service.PostRepository.FindById(ctx, id)
}

func (service *PostService) Create(ctx context.Context, post *model.Post) (*model.Post, error) {
	return service.PostRepository.Create(ctx, post)
}

func (service *PostService) Delete(ctx context.Context, id uint64) error {
	return service.PostRepository.Delete(ctx, id)
}

func (service *PostService) TotalRows(ctx context.Context) (int64, error) {
	return service.PostRepository.TotalRows(ctx)
}
