package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"helloadmin/api"
	"helloadmin/internal/ecode"
	"helloadmin/internal/model"
)

type RoleRepository interface {
	Find(ctx context.Context, request *api.RoleFindRequest) (int64, *[]model.Role, error)
	GetById(ctx context.Context, id int64) (*model.Role, error)
	Create(ctx context.Context, role *model.Role) error
	Update(ctx context.Context, id int64, role *model.Role) error
	UpdateRoleMenu(ctx context.Context, id int64, req *api.RoleMenuRequest) error
	Delete(ctx context.Context, id int64) error
	HasUser(ctx context.Context, id int64) int64
}

func NewRoleRepository(r *Repository) RoleRepository {
	return &roleRepository{
		Repository: r,
	}
}

type roleRepository struct {
	*Repository
}

func (r *roleRepository) Find(ctx context.Context, req *api.RoleFindRequest) (int64, *[]model.Role, error) {
	var count int64
	var role []model.Role
	query := r.DB(ctx)
	if req.Name != "" {
		query = query.Where("name = ?", req.Name)
	}
	if req.Page > 0 {
		query = query.Offset((req.Page - 1) * req.Size).Limit(req.Size)
	}
	query.Model(model.Role{}).Count(&count)
	if result := query.Order("id desc").Find(&role); result.Error != nil {
		return count, nil, result.Error
	}
	return count, &role, nil
}

func (r *roleRepository) GetById(ctx context.Context, id int64) (*model.Role, error) {
	var role model.Role
	if err := r.DB(ctx).Where("id = ?", id).Preload("Menus").First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ecode.ErrNotFound
		}
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) Create(ctx context.Context, role *model.Role) error {
	if err := r.DB(ctx).Create(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) Update(ctx context.Context, id int64, role *model.Role) error {
	if err := r.DB(ctx).Model(role).Where("id = ?", id).Updates(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) UpdateRoleMenu(ctx context.Context, id int64, req *api.RoleMenuRequest) error {
	var role model.Role
	if err := r.DB(ctx).Preload("Menus").First(&role, id).Error; err != nil {
		return err
	}
	// 清空角色原有的关联菜单
	if err := r.DB(ctx).Model(&role).Association("Menus").Clear(); err != nil {
		return err
	}
	var menus []model.Menu
	if err := r.DB(ctx).Find(&menus, req.MenuId).Error; err != nil {
		return err
	}
	if err := r.DB(ctx).Model(&role).Association("Menus").Append(menus); err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) Delete(ctx context.Context, id int64) error {
	if err := r.DB(ctx).Delete(&model.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) HasUser(ctx context.Context, id int64) int64 {
	var count int64
	r.DB(ctx).Model(model.User{}).Where("role_id = ?", id).Count(&count)
	return count
}
