package service

import (
	"context"

	"github.com/arvians-id/go-gorm/internal/model"
	"github.com/arvians-id/go-gorm/internal/repository"
)

type CommentService interface {
	List(ctx context.Context) ([]*model.Comment, error)
	FindById(ctx context.Context, id uint64) (*model.Comment, error)
	Create(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	Delete(ctx context.Context, id uint64) error
}

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
}

func NewCommentService(commentRepository repository.CommentRepository) CommentService {
	return &CommentServiceImpl{
		CommentRepository: commentRepository,
	}
}

func (service *CommentServiceImpl) List(ctx context.Context) ([]*model.Comment, error) {
	return service.CommentRepository.List(ctx)
}

func (service *CommentServiceImpl) FindById(ctx context.Context, id uint64) (*model.Comment, error) {
	return service.CommentRepository.FindById(ctx, id)
}

func (service *CommentServiceImpl) Create(ctx context.Context, post *model.Comment) (*model.Comment, error) {
	return service.CommentRepository.Create(ctx, post)
}

func (service *CommentServiceImpl) Delete(ctx context.Context, id uint64) error {
	return service.CommentRepository.Delete(ctx, id)
}
