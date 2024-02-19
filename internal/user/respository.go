package user

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"helloadmin/internal/ecode"
	"helloadmin/internal/repository"
)

type Repository interface {
	Create(ctx context.Context, user *Model) error
	Update(ctx context.Context, user *Model) error
	GetByID(ctx context.Context, id string) (*Model, error)
	GetByEmail(ctx context.Context, email string) (*Model, error)
	Search(ctx context.Context, request *FindRequest) (int64, *[]Model, error)
}

func NewRepository(r *repository.Repository) Repository {
	return &userRepository{
		Repository: r,
	}
}

type userRepository struct {
	*repository.Repository
}

func (r *userRepository) Create(ctx context.Context, user *Model) error {
	if err := r.DB(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *Model) error {
	if err := r.DB(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, userId string) (*Model, error) {
	var user Model
	if err := r.DB(ctx).Preload("Role").Preload("Department").Where("user_id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ecode.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*Model, error) {
	var user Model
	if err := r.DB(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Search(ctx context.Context, request *FindRequest) (int64, *[]Model, error) {
	var (
		users []Model
		total int64
	)
	query := r.DB(ctx)
	if request.Email != "" {
		query = query.Where("email = ?", request.Email)
	}
	if request.Nickname != "" {
		query = query.Where("nickname = ?", request.Nickname)
	}
	if request.RoleId != 0 {
		query = query.Where("role_id = ?", request.RoleId)
	}
	if request.DeptId != 0 {
		query = query.Where("dept_id = ?", request.DeptId)
	}
	if err := query.Model(Model{}).Count(&total).Error; err != nil {
		return 0, nil, err
	}
	if err := query.Order("id desc").Preload("Role").Preload("Department").Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&users).Error; err != nil {
		return total, nil, err
	}
	return total, &users, nil
}
