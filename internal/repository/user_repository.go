package repository

import (
	"context"
	"github.com/arvians-id/go-gorm/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	List(ctx context.Context) ([]*model.User, error)
	FindById(ctx context.Context, id uint64) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint64) error
	UpdateRoles(ctx context.Context, userId uint64, roles []*model.Role) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) List(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := repository.DB.WithContext(ctx).
		Table("user_roles").
		Select("COUNT(*) as total,*").
		Joins("LEFT JOIN roles r on r.id = user_roles.role_id").
		Joins("LEFT JOIN users u on u.id = user_roles.user_id").
		Group("u.id").
		Order("total desc").
		Preload("Roles").
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, id uint64) (*model.User, error) {
	var user *model.User
	err := repository.DB.WithContext(ctx).Preload("Roles").First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, user *model.User) (*model.User, error) {
	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Create(&user).Association("Roles").Append(&model.Role{ID: 2, Role: "member"})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, user *model.User) error {
	_, err := repository.FindById(ctx, user.ID)
	if err != nil {
		return err
	}

	err = repository.DB.WithContext(ctx).Updates(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	user, err := repository.FindById(ctx, id)
	if err != nil {
		return err
	}

	err = repository.DB.WithContext(ctx).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *UserRepositoryImpl) UpdateRoles(ctx context.Context, userId uint64, roles []*model.Role) error {
	user, err := repository.FindById(ctx, userId)
	if err != nil {
		return err
	}

	err = repository.DB.WithContext(ctx).Model(&user).Association("Roles").Replace(roles)
	if err != nil {
		return err
	}

	return nil
}
