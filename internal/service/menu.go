package service

import (
	"context"
	"helloadmin/api"
	"helloadmin/internal/model"
	"helloadmin/internal/repository"
	"time"
)

type MenuService interface {
	GetMenuById(ctx context.Context, id int64) (*api.MenuResponseItem, error)
	SearchMenu(ctx context.Context, request *api.MenuFindRequest) (*[]model.Menu, error)
	CreateMenu(ctx context.Context, request *api.MenuCreateRequest) error
	UpdateMenu(ctx context.Context, id int64, request *api.MenuUpdateRequest) error
	DeleteMenu(ctx context.Context, id int64) error
}

func NewMenuService(service *Service, menuRepository repository.MenuRepository) MenuService {
	return &menuService{
		Service:        service,
		menuRepository: menuRepository,
	}
}

type menuService struct {
	*Service
	menuRepository repository.MenuRepository
}

func (s *menuService) GetMenuById(ctx context.Context, id int64) (*api.MenuResponseItem, error) {
	menu, err := s.menuRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &api.MenuResponseItem{
		ID:        menu.ID,
		Name:      menu.Name,
		Title:     menu.Title,
		Icon:      menu.Icon,
		Path:      menu.Path,
		Type:      menu.Type,
		Action:    menu.Action,
		ParentId:  menu.ParentId,
		Component: menu.Component,
		Sort:      menu.Sort,
		Visible:   menu.Visible,
		CreatedAt: menu.CreatedAt.Format(time.RFC3339),
		UpdateAt:  menu.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *menuService) SearchMenu(ctx context.Context, req *api.MenuFindRequest) (*[]model.Menu, error) {
	return s.menuRepository.Find(ctx, req)
}

func (s *menuService) CreateMenu(ctx context.Context, req *api.MenuCreateRequest) error {
	menu := model.Menu{
		Name:      req.Name,
		Title:     req.Title,
		Icon:      req.Icon,
		Path:      req.Path,
		Type:      req.Type,
		Action:    req.Action,
		ParentId:  req.ParentId,
		Component: req.Component,
		Sort:      req.Sort,
		Visible:   req.Visible,
	}
	return s.menuRepository.Create(ctx, &menu)
}

func (s *menuService) UpdateMenu(ctx context.Context, id int64, req *api.MenuUpdateRequest) error {
	menu, err := s.menuRepository.GetById(ctx, id)
	if err != nil {
		return err
	}
	menu.Name = req.Name
	menu.Title = req.Title
	menu.Icon = req.Icon
	menu.Path = req.Path
	menu.Type = req.Type
	menu.Action = req.Action
	menu.ParentId = req.ParentId
	menu.Component = req.Component
	menu.Sort = req.Sort
	menu.Visible = req.Visible
	return s.menuRepository.Update(ctx, id, menu)
}

func (s *menuService) DeleteMenu(ctx context.Context, id int64) error {
	return s.menuRepository.Delete(ctx, id)
}
