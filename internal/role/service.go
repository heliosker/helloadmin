package role

import (
	"context"
	"helloadmin/internal/ecode"
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
	return &service{
		repo: repo,
	}
}

type service struct {
	repo Repository
}

func (s *service) GetRoleById(ctx context.Context, id int64) (*ResponseItem, error) {
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

func (s *service) SearchRole(ctx context.Context, req *FindRequest) (*Response, error) {
	var result Response
	count, roles, err := s.repo.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	result.Items = make([]ResponseItem, 0)
	if count > 0 {
		for _, role := range *roles {
			tmp := ResponseItem{
				Id:        role.ID,
				Name:      role.Name,
				Describe:  role.Describe,
				UpdatedAt: role.UpdatedAt.Format(time.DateTime),
				CreatedAt: role.CreatedAt.Format(time.DateTime),
			}
			result.Items = append(result.Items, tmp)
		}
	}
	return &result, nil
}

func (s *service) CreateRole(ctx context.Context, req *CreateRequest) error {
	role := Model{
		Name:     req.Name,
		Describe: req.Describe,
	}
	return s.repo.Create(ctx, &role)
}

func (s *service) UpdateRole(ctx context.Context, id int64, req *UpdateRequest) error {
	role := Model{
		Name:     req.Name,
		Describe: req.Describe,
	}
	return s.repo.Update(ctx, id, &role)
}

func (s *service) UpdateRoleMenu(ctx context.Context, id int64, req *MenuRequest) error {
	return s.repo.UpdateRoleMenu(ctx, id, req)
}

func (s *service) DeleteRole(ctx context.Context, id int64) error {
	if s.repo.HasUser(ctx, id) > 0 {
		return ecode.ErrRoleHasUser
	}
	return s.repo.Delete(ctx, id)
}
