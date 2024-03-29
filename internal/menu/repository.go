package menu

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"helloadmin/internal/ecode"
	"helloadmin/internal/repository"
)

type Repository interface {
	Find(ctx context.Context, request *FindRequest) (int64, *[]Model, error)
	FindByType(ctx context.Context, typ string) (*[]Model, error)
	GetById(ctx context.Context, id int64) (*Model, error)
	GetByParentId(ctx context.Context, id int64) (*[]Model, error)
	Create(ctx context.Context, menu *Model) error
	Update(ctx context.Context, id int64, menu *Model) error
	Delete(ctx context.Context, id int64) error
}

func NewRepository(r *repository.Repository) Repository {
	return &menuRepository{
		Repository: r,
	}
}

type menuRepository struct {
	*repository.Repository
}

func (r *menuRepository) FindByType(ctx context.Context, typ string) (*[]Model, error) {
	var menu []Model
	if err := r.DB(ctx).Where("type = ?", typ).Find(&menu).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepository) Find(ctx context.Context, req *FindRequest) (int64, *[]Model, error) {
	var count int64
	var menu []Model
	query := r.DB(ctx)
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Visible != "" {
		query = query.Where("visible = ?", req.Visible)
	}
	query.Model(&menu).Count(&count)
	if err := query.Order("sort asc").Find(&menu).Error; err != nil {
		return 0, nil, err
	}
	return count, &menu, nil
}

func (r *menuRepository) GetByParentId(ctx context.Context, id int64) (*[]Model, error) {
	var menu []Model
	query := r.DB(ctx)
	if err := query.Where("parent_id = ?", id).Find(&menu).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepository) GetById(ctx context.Context, id int64) (*Model, error) {
	var menu Model
	if err := r.DB(ctx).Where("id = ?", id).First(&menu).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ecode.ErrNotFound
		}
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepository) Create(ctx context.Context, menu *Model) error {
	if err := r.DB(ctx).Create(menu).Error; err != nil {
		return err
	}
	return nil
}

func (r *menuRepository) Update(ctx context.Context, id int64, menu *Model) error {
	menu.ID = uint(id)
	if err := r.DB(ctx).Model(menu).Omit("updated_at").Updates(menu).Error; err != nil {
		return err
	}
	return nil
}

func (r *menuRepository) Delete(ctx context.Context, id int64) error {
	if err := r.DB(ctx).Delete(&Model{}, id).Error; err != nil {
		return err
	}
	return nil
}
