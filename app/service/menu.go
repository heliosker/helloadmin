package service

import (
	"helloadmin/app/models"
	"helloadmin/pkg/app"
)

type MenuListReq struct {
	Name string `form:"name" binding:"max=100"`
}

func (svc *Service) GetMenuList(param *MenuListReq, meta app.Meta) ([]*models.Menu, error) {
	return svc.dao.GetTreeMenu(param.Name, meta.Page, meta.Size)
}

func (svc *Service) GetOptions() (map[uint]string, error) {
	return svc.dao.GetOptions()
}
