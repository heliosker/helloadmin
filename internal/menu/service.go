package menu

import (
	"context"
	"time"
)

type Service interface {
	GetMenuById(ctx context.Context, id int64) (*ResponseItem, error)
	SearchMenu(ctx context.Context, request *FindRequest) (*[]ResponseItem, error)
	CreateMenu(ctx context.Context, request *CreateRequest) error
	UpdateMenu(ctx context.Context, id int64, request *UpdateRequest) error
	DeleteMenu(ctx context.Context, id int64) error
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

type service struct {
	repo Repository
}

func (s *service) GetMenuById(ctx context.Context, id int64) (*ResponseItem, error) {
	menu, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &ResponseItem{
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

func (s *service) SearchMenu(ctx context.Context, req *FindRequest) (*[]ResponseItem, error) {
	menuList, err := s.repo.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	// Convert the flat menu list to a tree structure
	menuTree := buildMenuTree(menuList, 0)
	return &menuTree, nil
}

func buildMenuTree(menuList *[]Model, parentId uint) []ResponseItem {
	var result []ResponseItem
	for _, menuItem := range *menuList {
		if menuItem.ParentId == parentId {
			child := ResponseItem{
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

func (s *service) CreateMenu(ctx context.Context, req *CreateRequest) error {
	menu := Model{
		Name:      req.Name,
		Title:     req.Title,
		Icon:      req.Icon,
		Path:      req.Path,
		Type:      req.Type,
		ParentId:  req.ParentId,
		Component: req.Component,
		Sort:      req.Sort,
		Visible:   req.Visible,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return s.repo.Create(ctx, &menu)
}

func (s *service) UpdateMenu(ctx context.Context, id int64, req *UpdateRequest) error {
	menu, err := s.repo.GetById(ctx, id)
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
	menu.UpdatedAt = time.Now()
	return s.repo.Update(ctx, id, menu)
}

func (s *service) DeleteMenu(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
