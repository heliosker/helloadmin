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
	SearchMenu(ctx context.Context, request *api.MenuFindRequest) (*[]api.MenuResponseItem, error)
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
		ParentId:  menu.ParentId,
		Component: menu.Component,
		Sort:      menu.Sort,
		Visible:   menu.Visible,
		CreatedAt: menu.CreatedAt.Format(time.RFC3339),
		UpdatedAt: menu.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *menuService) SearchMenu(ctx context.Context, req *api.MenuFindRequest) (*[]api.MenuResponseItem, error) {
	menuList, err := s.menuRepository.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	// Convert the flat menu list to a tree structure
	menuTree := buildMenuTree(menuList, 0)
	return &menuTree, nil
}

func buildMenuTree(menuList *[]model.Menu, parentId uint) []api.MenuResponseItem {
	var result []api.MenuResponseItem

	for _, menuItem := range *menuList {
		if menuItem.ParentId == parentId {
			child := api.MenuResponseItem{
				ID:        menuItem.ID,
				Name:      menuItem.Name,
				Title:     menuItem.Title,
				Icon:      menuItem.Icon,
				Path:      menuItem.Path,
				Type:      menuItem.Type,
				ParentId:  menuItem.ParentId,
				Component: menuItem.Component,
				Sort:      menuItem.Sort,
				Visible:   menuItem.Visible,
				CreatedAt: menuItem.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: menuItem.UpdatedAt.Format("2006-01-02 15:04:05"),
			}

			// Recursively build the tree for child items
			child.Children = buildMenuTree(menuList, menuItem.ID)
			result = append(result, child)
		}
	}

	return result
}

func (s *menuService) CreateMenu(ctx context.Context, req *api.MenuCreateRequest) error {
	menu := model.Menu{
		Name:      req.Name,
		Title:     req.Title,
		Icon:      req.Icon,
		Path:      req.Path,
		Type:      req.Type,
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
	menu.ParentId = req.ParentId
	menu.Component = req.Component
	menu.Sort = req.Sort
	menu.Visible = req.Visible
	return s.menuRepository.Update(ctx, id, menu)
}

func (s *menuService) DeleteMenu(ctx context.Context, id int64) error {
	return s.menuRepository.Delete(ctx, id)
}
