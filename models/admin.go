package models

import "gorm.io/gorm"

type AdminUser struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func (AdminUser) TableName() string {
	return "hi_admin_users"
}

func AdminUserExist(username, password string) bool {
	var admin AdminUser
	DB.Select("id").Where(AdminUser{Username: username, Password: password}).First(&admin)
	if admin.ID > 0 {
		return true
	}
	return false
}
