package role

import (
	"context"
	"time"
)

type Service interface {
	GetRoleById(ctx context.Context, id int64) (*ResponseItem, error)
	SearchRole(ctx context.Context, request *FindRequest) (*Response, error)
	CreateRole(ctx context.Context, request *CreateRequest) error
	UpdateRole(ctx context.Context, id int64, request *UpdateRequest) error
	DeleteRole(ctx context.Context, id int64) error
	UpdateRoleMenu(ctx context.Context, id int64, request *MenuRequest) error
}

func NewService(repo Repository) Service {
	return &roleService{
		repo: repo,
	}
}

type roleService struct {
	repo Repository
}

func (s *roleService) GetRoleById(ctx context.Context, id int64) (*ResponseItem, error) {
	role, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	menuIds := make([]uint, 0)
	if len(role.Menus) > 0 {
		for _, menu := range role.Menus {
			menuIds = append(menuIds, menu.ID)
		}
	}
	return &ResponseItem{
		Id:        role.ID,
		Name:      role.Name,
		Describe:  role.Describe,
		UpdatedAt: role.UpdatedAt.Format(time.DateTime),
		CreatedAt: role.CreatedAt.Format(time.DateTime),
		MenuId:    menuIds,
	}, nil
}

func (s *roleService) SearchRole(ctx context.Context, req *FindRequest) (*Response, error) {
	var result Response
	count, roles, err := s.repo.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	result.Items = make([]ResponseItem, 0)
	if count > 0 {
		for _, role := range *roles {
			menuIds := make([]uint, 0)
			if len(role.Menus) > 0 {
				for _, menu := range role.Menus {
					menuIds = append(menuIds, menu.ID)
				}
			}
			tmp := ResponseItem{
				Id:        role.ID,
				Name:      role.Name,
				Describe:  role.Describe,
				MenuId:    menuIds,
				UpdatedAt: role.UpdatedAt.Format(time.DateTime),
				CreatedAt: role.CreatedAt.Format(time.DateTime),
			}
			result.Items = append(result.Items, tmp)
		}
	}
	return &result, nil
}

func (s *roleService) CreateRole(ctx context.Context, req *CreateRequest) error {
	role := Model{
		Name:      req.Name,
		Describe:  req.Describe,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return s.repo.Create(ctx, &role)
}

func (s *roleService) UpdateRole(ctx context.Context, id int64, req *UpdateRequest) error {
	role := Model{
		Name:      req.Name,
		Describe:  req.Describe,
		UpdatedAt: time.Now(),
	}
	return s.repo.Update(ctx, id, &role)
}

func (s *roleService) UpdateRoleMenu(ctx context.Context, id int64, req *MenuRequest) error {
	return s.repo.UpdateRoleMenu(ctx, id, req)
}

func (s *roleService) DeleteRole(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
