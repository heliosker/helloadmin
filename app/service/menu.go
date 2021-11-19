package service

import (
	"helloadmin/app/models"
)

type MenuListReq struct {
	Name string `form:"name" binding:"max=100"`
}

type MenuTreeMap struct {
	models.Menu
	Children []models.Menu `json:"children" gorm:"-"`
}

func (svc *Service) GetTreeMenu(param *MenuListReq) ([]MenuTreeMap, error) {

	var menuTree = make([]MenuTreeMap, svc.dao.Count(0))
	menus, e := svc.dao.GetTreeMenu(param.Name)
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
			menuTree[k].Children = svc.GetChildren(v.ID, menus)
		}
	}
	return menuTree, e
}

func (svc *Service) GetChildren(id uint, menu []*models.Menu) []models.Menu {
	menus := []models.Menu{}
	for _, v := range menu {
		if v.ParentId == id {
			menus = append(menus, *v)
		}
	}
	return menus
}

func (svc *Service) GetOptions() ([]map[string]interface{}, error) {
	return svc.dao.GetOptions()
}
