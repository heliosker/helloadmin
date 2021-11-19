package service

import (
	"helloadmin/app/models"
)

type MenuListReq struct {
	Name string `form:"name" binding:"max=100"`
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
			menuTree[k].Title = v.Title
			menuTree[k].Sort = v.Sort
			menuTree[k].Icon = v.Icon
			menuTree[k].Uri = v.Uri
			menuTree[k].Extension = v.Extension
			menuTree[k].IsShow = v.IsShow
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
