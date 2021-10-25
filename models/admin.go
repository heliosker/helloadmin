package models

import (
	"gorm.io/gorm"
)

type AdminUser struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (AdminUser) TableName() string {
	return "hi_admin_users"
}

func AdminUserExist(email, password string) bool {
	var admin AdminUser
	DB.Select("id").Where(AdminUser{Email: email, Password: password}).First(&admin)
	if admin.ID > 0 {
		return true
	}
	return false
}
