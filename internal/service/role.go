package service

import (
	"context"
	"helloadmin/api"
	"helloadmin/internal/ecode"
	"helloadmin/internal/model"
	"helloadmin/internal/repository"
	"time"
)

type RoleService interface {
	GetRoleById(ctx context.Context, id int64) (*api.RoleResponseItem, error)
	SearchRole(ctx context.Context, request *api.RoleFindRequest) (*api.RoleResponse, error)
	CreateRole(ctx context.Context, request *api.RoleCreateRequest) error
	UpdateRole(ctx context.Context, id int64, request *api.RoleUpdateRequest) error
	DeleteRole(ctx context.Context, id int64) error
}

func NewRoleService(service *Service, roleRepository repository.RoleRepository) RoleService {
	return &roleService{
		Service:        service,
		roleRepository: roleRepository,
	}
}

type roleService struct {
	*Service
	roleRepository repository.RoleRepository
}

func (s *roleService) GetRoleById(ctx context.Context, id int64) (*api.RoleResponseItem, error) {
	role, err := s.roleRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &api.RoleResponseItem{
		Id:        role.ID,
		Name:      role.Name,
		Slug:      role.Slug,
		Describe:  role.Describe,
		UpdatedAt: role.UpdatedAt.Format(time.DateTime),
		CreatedAt: role.CreatedAt.Format(time.DateTime),
	}, nil
}

func (s *roleService) SearchRole(ctx context.Context, req *api.RoleFindRequest) (*api.RoleResponse, error) {
	var result api.RoleResponse
	count, roles, err := s.roleRepository.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	result.Items = make([]api.RoleResponseItem, 0)
	if count > 0 {
		for _, role := range *roles {
			tmp := api.RoleResponseItem{
				Id:        role.ID,
				Name:      role.Name,
				Slug:      role.Slug,
				Describe:  role.Describe,
				UpdatedAt: role.UpdatedAt.Format(time.DateTime),
				CreatedAt: role.CreatedAt.Format(time.DateTime),
			}
			result.Items = append(result.Items, tmp)
		}
	}
	result.Pagination = api.Pagination{
		Page:  req.Page,
		Size:  req.Size,
		Count: int(count),
	}
	return &result, nil
}

func (s *roleService) CreateRole(ctx context.Context, req *api.RoleCreateRequest) error {
	role := model.Role{
		Name:     req.Name,
		Slug:     req.Slug,
		Describe: req.Describe,
	}
	return s.roleRepository.Create(ctx, &role)
}

func (s *roleService) UpdateRole(ctx context.Context, id int64, req *api.RoleUpdateRequest) error {
	role := model.Role{
		Name:     req.Name,
		Slug:     req.Slug,
		Describe: req.Describe,
	}
	return s.roleRepository.Update(ctx, id, &role)
}

func (s *roleService) DeleteRole(ctx context.Context, id int64) error {
	if s.roleRepository.HasUser(ctx, id) > 0 {
		return ecode.ErrRoleHasUser
	}
	return s.roleRepository.Delete(ctx, id)
}
