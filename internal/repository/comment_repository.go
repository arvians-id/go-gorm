package repository

import (
	"context"
	"github.com/arvians-id/go-gorm/internal/model"
	"gorm.io/gorm"
)

type CommentRepository interface {
	List(ctx context.Context) ([]*model.Comment, error)
	FindById(ctx context.Context, id uint64) (*model.Comment, error)
	Create(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	Delete(ctx context.Context, id uint64) error
}

type CommentRepositoryImpl struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &CommentRepositoryImpl{
		DB: db,
	}
}

func (repository *CommentRepositoryImpl) List(ctx context.Context) ([]*model.Comment, error) {
	var comments []*model.Comment
	err := repository.DB.WithContext(ctx).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (repository *CommentRepositoryImpl) FindById(ctx context.Context, id uint64) (*model.Comment, error) {
	var comment *model.Comment
	err := repository.DB.WithContext(ctx).First(&comment, id).Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (repository *CommentRepositoryImpl) Create(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	err := repository.DB.WithContext(ctx).Create(&comment).Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (repository *CommentRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	comment, err := repository.FindById(ctx, id)
	if err != nil {
		return err
	}

	err = repository.DB.WithContext(ctx).Delete(&comment).Error
	if err != nil {
		return err
	}

	return nil
}
