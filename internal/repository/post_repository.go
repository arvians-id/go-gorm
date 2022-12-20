package repository

import (
	"context"
	"github.com/arvians-id/go-gorm/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostRepositoryContract interface {
	List(ctx context.Context, perPage int, offset int) ([]*model.Post, error)
	FindById(ctx context.Context, id uint64) (*model.PostResponse, error)
	Create(ctx context.Context, post *model.Post) (*model.Post, error)
	Delete(ctx context.Context, id uint64) error
	TotalRows(ctx context.Context) (int64, error)
}

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return PostRepository{
		DB: db,
	}
}

func (repository *PostRepository) List(ctx context.Context, perPage int, offset int) ([]*model.Post, error) {
	var posts []*model.Post
	err := repository.DB.WithContext(ctx).Preload(clause.Associations, func(db *gorm.DB) *gorm.DB {
		return db.Order("comments.created_at desc")
	}).Order("posts.created_at desc").Limit(perPage).Offset(offset).Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (repository *PostRepository) FindById(ctx context.Context, id uint64) (*model.PostResponse, error) {
	var post *model.PostResponse
	err := repository.DB.WithContext(ctx).Model(&model.Post{}).
		Select("posts.*, comments.body as comment_body").
		Where("posts.id = ?", id).
		Joins("Comments").
		First(&post).Error
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (repository *PostRepository) Create(ctx context.Context, post *model.Post) (*model.Post, error) {
	err := repository.DB.WithContext(ctx).Create(&post).Error
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (repository *PostRepository) Delete(ctx context.Context, id uint64) error {
	_, err := repository.FindById(ctx, id)
	if err != nil {
		return err
	}

	var post model.Post
	err = repository.DB.WithContext(ctx).Where("id = ?", id).Delete(&post).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *PostRepository) TotalRows(ctx context.Context) (int64, error) {
	var totalRows int64
	err := repository.DB.WithContext(ctx).Model(&model.Post{}).Count(&totalRows).Error
	if err != nil {
		return 0, err
	}

	return totalRows, nil
}
