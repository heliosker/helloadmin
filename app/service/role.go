package service

import (
	"helloadmin/app/models"
	"helloadmin/pkg/convert"
	"strings"
)

type CreateRole struct {
	Name     string `form:"name" binding:"required"`
	Describe string `form:"describe"`
}

type UpdateRole struct {
	Name     string `form:"name" binding:"required"`
	Describe string `form:"describe"`
}

type CreateRoleMenu struct {
	MenuId string `form:"menu_id" binding:"required"`
}

type ResultRole struct {
	Role models.Role
	Menu []models.RoleMenu
}

func (svc *Service) CreateRole(param CreateRole) error {
	return svc.dao.CreateRole(param.Name, param.Describe)
}

func (svc *Service) UpdateRole(id uint, param UpdateRole) error {
	return svc.dao.UpdateRole(id, param.Name, param.Describe)
}

func (svc *Service) DeleteRole(roleId uint) error {
	return svc.dao.DeleteRole(roleId)
}

func (svc *Service) FindRole(roleId uint) ResultRole {
	var retRole ResultRole
	role := svc.dao.GetRole(roleId)
	roleMenu := svc.dao.GetRoleMenu(roleId)
	retRole.Role = role
	retRole.Menu = roleMenu
	return retRole
}

// 新增角色菜单
func (svc *Service) CreateRoleMenu(roleId uint, param CreateRoleMenu) error {
	if strings.Contains(param.MenuId, ",") {
		for _, v := range strings.Split(param.MenuId, ",") {
			if err := svc.dao.CreateRoleMenu(roleId, convert.StrTo(v).MustUInt()); err != nil {
				return err
			}
		}
	} else {
		return svc.dao.CreateRoleMenu(roleId, convert.StrTo(param.MenuId).MustUInt())
	}
	return nil
}

// 删除角色菜单
func (svc *Service) DeleteRoleMenu(roleId uint) error {
	return svc.dao.DeleteRoleMenu(roleId)
}
