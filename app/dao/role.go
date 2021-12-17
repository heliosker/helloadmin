package dao

import "helloadmin/app/models"

func (d *Dao) CreateRole(name, describe string) error {
	role := models.Role{
		Name:     name,
		Describe: describe,
	}
	return role.Create(d.engine)
}

func (d *Dao) GetRole(roleId uint) models.Role {
	role := models.Role{}
	role.ID = roleId
	return role.One(d.engine)
}

func (d *Dao) UpdateRole(id uint, name, describe string) error {
	role := models.Role{Name: name, Describe: describe}
	role.ID = id
	return role.Update(d.engine)
}

func (d *Dao) DeleteRole(roleId uint) error {
	role := models.Role{}
	role.ID = roleId
	return role.Delete(d.engine)
}

func (d *Dao) GetRoleMenu(roleId uint) []models.RoleMenu {
	role := models.RoleMenu{}
	role.RoleId = roleId
	return role.List(d.engine)
}

func (d *Dao) CreateRoleMenu(roleId, menuId uint) error {
	rm := models.RoleMenu{RoleId: roleId, MenuId: menuId}
	return rm.Create(d.engine)
}

func (d *Dao) DeleteRoleMenu(roleId uint) error {
	rm := models.RoleMenu{RoleId: roleId}
	return rm.Delete(d.engine)
}
