package service

import (
	"context"

	"github.com/arvians-id/go-gorm/internal/model"
	"github.com/arvians-id/go-gorm/internal/repository"
)

type CommentServiceContract interface {
	List(ctx context.Context) ([]*model.Comment, error)
	FindById(ctx context.Context, id uint64) (*model.Comment, error)
	Create(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	Delete(ctx context.Context, id uint64) error
}

type CommentService struct {
	CommentRepository repository.CommentRepository
}

func NewCommentService(commentRepository repository.CommentRepository) CommentService {
	return CommentService{
		CommentRepository: commentRepository,
	}
}

func (service *CommentService) List(ctx context.Context) ([]*model.Comment, error) {
	return service.CommentRepository.List(ctx)
}

func (service *CommentService) FindById(ctx context.Context, id uint64) (*model.Comment, error) {
	return service.CommentRepository.FindById(ctx, id)
}

func (service *CommentService) Create(ctx context.Context, post *model.Comment) (*model.Comment, error) {
	return service.CommentRepository.Create(ctx, post)
}

func (service *CommentService) Delete(ctx context.Context, id uint64) error {
	return service.CommentRepository.Delete(ctx, id)
}
