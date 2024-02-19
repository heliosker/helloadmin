package department

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"helloadmin/internal/ecode"
	"helloadmin/internal/repository"
)

type Repository interface {
	Find(ctx context.Context, request *FindRequest) (int64, *[]Model, error)
	GetById(ctx context.Context, id int64) (*Model, error)
	GetByParentId(ctx context.Context, id int64) (*[]Model, error)
	Create(ctx context.Context, department *Model) error
	Update(ctx context.Context, id int64, department *Model) error
	Delete(ctx context.Context, id int64) error
}

func NewRepository(r *repository.Repository) Repository {
	return &departmentRepository{
		Repository: r,
	}
}

type departmentRepository struct {
	*repository.Repository
}

func (r *departmentRepository) Find(ctx context.Context, req *FindRequest) (int64, *[]Model, error) {
	var count int64
	var departments []Model
	query := r.DB(ctx)
	if req.Name != "" {
		query = query.Where("name = ?", req.Name)
	}
	query.Model(Model{}).Count(&count)
	if err := query.Order("sort DESC").Find(&departments).Error; err != nil {
		return count, nil, err
	}
	return count, &departments, nil
}

func (r *departmentRepository) GetById(ctx context.Context, id int64) (*Model, error) {
	var department Model
	if err := r.DB(ctx).Where("id = ?", id).First(&department).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ecode.ErrNotFound
		}
		return nil, err
	}
	return &department, nil
}

func (r *departmentRepository) GetByParentId(ctx context.Context, id int64) (*[]Model, error) {
	var departments []Model
	if err := r.DB(ctx).Where("parent_id = ?", id).Find(&departments).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &departments, nil
}

func (r *departmentRepository) Create(ctx context.Context, department *Model) error {
	if err := r.DB(ctx).Create(department).Error; err != nil {
		return err
	}
	return nil
}

func (r *departmentRepository) Update(ctx context.Context, id int64, department *Model) error {
	if err := r.DB(ctx).Model(department).Select("name", "parent_id", "sort", "leader").Where("id = ?", id).Updates(department).Error; err != nil {
		return err
	}
	return nil
}

func (r *departmentRepository) Delete(ctx context.Context, id int64) error {
	var users []Model
	if err := r.DB(ctx).Where("dept_id = ?", id).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	if len(users) > 0 {
		return ecode.ErrDeptHasUser
	}
	if err := r.DB(ctx).Delete(&Model{}, id).Error; err != nil {
		return err
	}
	return nil
}
