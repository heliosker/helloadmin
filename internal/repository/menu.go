package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"helloadmin/api"
	"helloadmin/internal/ecode"
	"helloadmin/internal/model"
)

type MenuRepository interface {
	Find(ctx context.Context, request *api.MenuFindRequest) (*[]model.Menu, error)
	GetById(ctx context.Context, id int64) (*model.Menu, error)
	Create(ctx context.Context, menu *model.Menu) error
	Update(ctx context.Context, id int64, menu *model.Menu) error
	Delete(ctx context.Context, id int64) error
}

func NewMenuRepository(r *Repository) MenuRepository {
	return &menuRepository{
		Repository: r,
	}
}

type menuRepository struct {
	*Repository
}

func (r *menuRepository) Find(ctx context.Context, req *api.MenuFindRequest) (*[]model.Menu, error) {
	var count int64
	var menu []model.Menu
	query := r.DB(ctx)
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Visible != "" {
		query = query.Where("visible = ?", req.Visible)
	}
	if req.Page > 0 {
		query = query.Offset((req.Page - 1) * req.Size).Limit(req.Size)
	}
	query.Model(&menu).Count(&count)
	if err := query.Order("sort desc").Find(&menu).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepository) GetById(ctx context.Context, id int64) (*model.Menu, error) {
	var menu model.Menu
	if err := r.DB(ctx).Where("id = ?", id).First(&menu).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ecode.ErrNotFound
		}
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepository) Create(ctx context.Context, menu *model.Menu) error {
	if err := r.DB(ctx).Create(menu).Error; err != nil {
		return err
	}
	return nil
}

func (r *menuRepository) Update(ctx context.Context, id int64, menu *model.Menu) error {
	menu.ID = uint(id)
	if err := r.DB(ctx).Model(menu).Omit("updated_at").Updates(menu).Error; err != nil {
		return err
	}
	return nil
}

func (r *menuRepository) Delete(ctx context.Context, id int64) error {
	if err := r.DB(ctx).Delete(&model.Menu{}, id).Error; err != nil {
		return err
	}
	return nil
}
