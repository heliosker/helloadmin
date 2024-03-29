package menu

import (
	"context"
	"time"

	"helloadmin/internal/ecode"
)

type Service interface {
	GetMenuById(ctx context.Context, id int64) (*ResponseItem, error)
	SearchMenu(ctx context.Context, request *FindRequest) (Response, error)
	CreateMenu(ctx context.Context, request *CreateRequest) error
	UpdateMenu(ctx context.Context, id int64, request *UpdateRequest) error
	DeleteMenu(ctx context.Context, id int64) error
	Options(ctx context.Context, req *OptionRequest) ([]Option, error)
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
		CreatedAt: menu.CreatedAt.Format(time.DateTime),
		UpdatedAt: menu.UpdatedAt.Format(time.DateTime),
	}, nil
}

func (s *service) SearchMenu(ctx context.Context, req *FindRequest) (resp Response, err error) {
	resp.Items = make([]ResponseItem, 0)
	count, menuList, err := s.repo.Find(ctx, req)
	if err != nil {
		return resp, err
	}
	if count > 0 {
		// Convert the flat menu list to a tree structure
		resp.Items = buildMenuTree(menuList, 0)
	}
	return resp, nil
}

func (s *service) Options(ctx context.Context, req *OptionRequest) ([]Option, error) {
	var menus *[]Model
	var err error
	switch req.Type {
	case TypeDirectory:
		return []Option{{Label: "顶级", Value: 0}}, nil
	case TypeMenu:
		menus, err = s.repo.FindByType(ctx, TypeDirectory)
		return buildOptions(*menus, true)
	case TypeButton:
		menus, err = s.repo.FindByType(ctx, TypeMenu)
		return buildOptions(*menus, false)
	default:
		return nil, err
	}
}

func buildOptions(menus []Model, top bool) ([]Option, error) {
	options := make([]Option, 1)
	if top {
		options[0] = Option{Label: "顶级", Value: 0}
	}
	if len(menus) > 0 {
		for _, item := range menus {
			options = append(options, Option{
				Label: item.Title,
				Value: item.ID,
			})
		}
		return options, nil
	}
	return options, nil
}

func buildMenuTree(menuList *[]Model, parentId uint) []ResponseItem {
	result := make([]ResponseItem, 0)
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
				CreatedAt: menuItem.CreatedAt.Format(time.DateTime),
				UpdatedAt: menuItem.UpdatedAt.Format(time.DateTime),
			}
			// Recursively build the tree for child items
			child.Children = buildMenuTree(menuList, menuItem.ID)
			result = append(result, child)
		}
	}
	return result
}

func (s *service) CreateMenu(ctx context.Context, req *CreateRequest) error {
	if req.ParentId > 0 {
		if menu, _ := s.repo.GetById(ctx, int64(req.ParentId)); menu == nil {
			return ecode.ErrMenuParentedNotFound
		}
	}
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
	menu, err := s.repo.GetByParentId(ctx, id)
	if err != nil {
		return err
	}
	// 删除菜单时，判断是否存在下级
	if len(*menu) > 0 {
		return ecode.ErrMenuHasChild
	}
	return s.repo.Delete(ctx, id)
}
