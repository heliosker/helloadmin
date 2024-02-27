package role

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"helloadmin/internal/ecode"
	"helloadmin/internal/menu"
	"helloadmin/internal/repository"
)

type Repository interface {
	Find(ctx context.Context, request *FindRequest) (int64, *[]Model, error)
	GetById(ctx context.Context, id int64) (*Model, error)
	Create(ctx context.Context, role *Model) error
	Update(ctx context.Context, id int64, role *Model) error
	UpdateRoleMenu(ctx context.Context, id int64, req *MenuRequest) error
	Delete(ctx context.Context, id int64) error
}

func NewRepository(r *repository.Repository) Repository {
	return &roleRepository{
		Repository: r,
	}
}

type roleRepository struct {
	*repository.Repository
}

func (r *roleRepository) Find(ctx context.Context, req *FindRequest) (int64, *[]Model, error) {
	var count int64
	var role []Model
	query := r.DB(ctx)
	if req.Name != "" {
		query = query.Where("name = ?", req.Name)
	}
	query.Model(Model{}).Count(&count)
	if result := query.Order("id desc").Preload("Menus").Find(&role); result.Error != nil {
		return count, nil, result.Error
	}
	return count, &role, nil
}

func (r *roleRepository) GetById(ctx context.Context, id int64) (*Model, error) {
	var role Model
	if err := r.DB(ctx).Where("id = ?", id).Preload("Menus").First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ecode.ErrRoleNotFound
		}
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) Create(ctx context.Context, role *Model) error {
	if err := r.DB(ctx).Create(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) Update(ctx context.Context, id int64, role *Model) error {
	if err := r.DB(ctx).Model(role).Where("id = ?", id).Updates(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) UpdateRoleMenu(ctx context.Context, id int64, req *MenuRequest) error {
	// 开启数据库事务
	tx := r.DB(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 回滚事务
		}
	}()

	var role Model
	if err := tx.Preload("Menus").First(&role, id).Error; err != nil {
		tx.Rollback() // 回滚事务
		return err
	}

	// 清空角色原有的关联菜单
	if err := tx.Model(&role).Association("Menus").Clear(); err != nil {
		tx.Rollback() // 回滚事务
		return err
	}

	if len(req.MenuId) > 0 {
		var menus []menu.Model
		if err := tx.Where("id IN (?)", req.MenuId).Find(&menus).Error; err != nil {
			tx.Rollback() // 回滚事务
			return err
		}

		// 将新的菜单关联到角色
		if err := tx.Model(&role).Association("Menus").Append(menus); err != nil {
			tx.Rollback() // 回滚事务
			return err
		}
	}

	// 提交事务
	return tx.Commit().Error
}

func (r *roleRepository) Delete(ctx context.Context, id int64) error {
	var count int64
	r.DB(ctx).Model(Model{}).Where("role_id = ?", id).Count(&count)
	if count > 0 {
		return ecode.ErrRoleHasUser
	}
	if err := r.DB(ctx).Delete(&Model{}, id).Error; err != nil {
		return err
	}
	return nil
}
