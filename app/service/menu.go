package service

import (
	"helloadmin/app/models"
)

type MenuListReq struct {
	Label string `form:"label" binding:"max=100"`
}

type MenuTreeMap struct {
	models.Menu
	Children []*models.Menu `json:"children" gorm:"-"`
}

func (svc *Service) GetTreeMenu(param *MenuListReq) ([]MenuTreeMap, error) {
	menus, e := svc.dao.GetChildren(0)
	var menuTree = make([]MenuTreeMap, len(menus))
	if e != nil {
		return menuTree, e
	}
	for k, v := range menus {
		if v.ParentId == 0 {
			menuTree[k].ID = v.ID
			menuTree[k].ParentId = v.ParentId
			menuTree[k].Label = v.Label
			menuTree[k].Sort = v.Sort
			menuTree[k].Icon = v.Icon
			menuTree[k].Path = v.Path
			menuTree[k].Redirect = v.Redirect
			menuTree[k].Show = v.Show
			if children, e := svc.dao.GetChildren(v.ID); e != nil {
			} else {
				menuTree[k].Children = children
			}
		}
	}
	return menuTree, e
}

func (svc *Service) GetOptions() ([]map[string]interface{}, error) {
	return svc.dao.GetOptions()
}
