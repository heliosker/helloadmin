package models

import "gorm.io/gorm"

type RoleMenu struct {
	ID     uint `json:"id",gorm:"primarykey"`
	RoleId uint `json:"role_id"`
	MenuId uint `json:"menu_id"`
}

func (RoleMenu) TableName() string {
	return "hi_role_menu"
}

func (rm RoleMenu) Create(db *gorm.DB) error {
	return db.Create(&rm).Error
}

func (rm RoleMenu) List(db *gorm.DB) []RoleMenu {
	var menus []RoleMenu
	db.Debug().Where("role_id = ?", rm.RoleId).Find(&menus)
	return menus
}

func (rm RoleMenu) Delete(db *gorm.DB) error {
	return db.Where("role_id = ?", rm.RoleId).Delete(&rm).Error
}
