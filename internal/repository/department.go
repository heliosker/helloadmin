package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"helloadmin/api"
	"helloadmin/internal/ecode"
	"helloadmin/internal/model"
)

type DepartmentRepository interface {
	Find(ctx context.Context, request *api.DepartmentFindRequest) (int64, *[]model.Department, error)
	GetById(ctx context.Context, id int64) (*model.Department, error)
	Create(ctx context.Context, department *model.Department) error
	Update(ctx context.Context, id int64, department *model.Department) error
	Delete(ctx context.Context, id int64) error
}

func NewDepartmentRepository(r *Repository) DepartmentRepository {
	return &departmentRepository{
		Repository: r,
	}
}

type departmentRepository struct {
	*Repository
}

func (r *departmentRepository) Find(ctx context.Context, req *api.DepartmentFindRequest) (int64, *[]model.Department, error) {
	var count int64
	var departments []model.Department
	query := r.DB(ctx)
	if req.Name != "" {
		query = query.Where("name = ?", req.Name)
	}
	if req.Page != 0 && req.Size != 0 {
		query = query.Offset((req.Page - 1) * req.Size).Limit(req.Size)
	}
	query.Model(model.Department{}).Count(&count)
	if err := query.Order("sort DESC").Find(&departments).Error; err != nil {
		return count, nil, err
	}
	return count, &departments, nil
}

func (r *departmentRepository) GetById(ctx context.Context, id int64) (*model.Department, error) {
	var department model.Department
	if err := r.DB(ctx).Where("id = ?", id).First(&department).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ecode.ErrNotFound
		}
		return nil, err
	}
	return &department, nil
}

func (r *departmentRepository) Create(ctx context.Context, department *model.Department) error {
	if err := r.DB(ctx).Create(department).Error; err != nil {
		return err
	}
	return nil
}

func (r *departmentRepository) Update(ctx context.Context, id int64, department *model.Department) error {
	if err := r.DB(ctx).Model(department).Select("name", "parent_id", "sort", "leader").Where("id = ?", id).Updates(department).Error; err != nil {
		return err
	}
	return nil
}

func (r *departmentRepository) Delete(ctx context.Context, id int64) error {
	if err := r.DB(ctx).Delete(&model.Department{}, id).Error; err != nil {
		return err
	}
	return nil
}
