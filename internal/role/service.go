package role

import (
	"context"
	"helloadmin/internal/api"
	"helloadmin/internal/ecode"
	"time"
)

type RoleService interface {
	GetRoleById(ctx context.Context, id int64) (*RoleResponseItem, error)
	SearchRole(ctx context.Context, request *RoleFindRequest) (*RoleResponse, error)
	CreateRole(ctx context.Context, request *RoleCreateRequest) error
	UpdateRole(ctx context.Context, id int64, request *RoleUpdateRequest) error
	DeleteRole(ctx context.Context, id int64) error
	UpdateRoleMenu(ctx context.Context, id int64, request *RoleMenuRequest) error
}

func NewRoleService(repo RoleRepository) RoleService {
	return &roleService{
		roleRepository: repo,
	}
}

type roleService struct {
	roleRepository RoleRepository
}

func (s *roleService) GetRoleById(ctx context.Context, id int64) (*RoleResponseItem, error) {
	role, err := s.roleRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	menuIds := make([]uint, 0)
	if len(role.Menus) > 0 {
		for _, menu := range role.Menus {
			menuIds = append(menuIds, menu.ID)
		}
	}
	return &RoleResponseItem{
		Id:        role.ID,
		Name:      role.Name,
		Describe:  role.Describe,
		UpdatedAt: role.UpdatedAt.Format(time.DateTime),
		CreatedAt: role.CreatedAt.Format(time.DateTime),
		MenuId:    menuIds,
	}, nil
}

func (s *roleService) SearchRole(ctx context.Context, req *RoleFindRequest) (*RoleResponse, error) {
	var result RoleResponse
	count, roles, err := s.roleRepository.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	result.Items = make([]RoleResponseItem, 0)
	if count > 0 {
		for _, role := range *roles {
			tmp := RoleResponseItem{
				Id:        role.ID,
				Name:      role.Name,
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

func (s *roleService) CreateRole(ctx context.Context, req *RoleCreateRequest) error {
	role := Model{
		Name:     req.Name,
		Describe: req.Describe,
	}
	return s.roleRepository.Create(ctx, &role)
}

func (s *roleService) UpdateRole(ctx context.Context, id int64, req *RoleUpdateRequest) error {
	role := Model{
		Name:     req.Name,
		Describe: req.Describe,
	}
	return s.roleRepository.Update(ctx, id, &role)
}

func (s *roleService) UpdateRoleMenu(ctx context.Context, id int64, req *RoleMenuRequest) error {
	return s.roleRepository.UpdateRoleMenu(ctx, id, req)
}

func (s *roleService) DeleteRole(ctx context.Context, id int64) error {
	if s.roleRepository.HasUser(ctx, id) > 0 {
		return ecode.ErrRoleHasUser
	}
	return s.roleRepository.Delete(ctx, id)
}
