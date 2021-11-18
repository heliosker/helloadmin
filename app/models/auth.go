package models

import (
	"gorm.io/gorm"
	"helloadmin/pkg/errcode"
)

type Auth struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}

func (a Auth) TableName() string {
	return "hi_admin_users"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var admin Auth
	err := db.Debug().Where(AdminUser{Username: a.Username, Password: a.Password}).First(&admin).Error
	if admin.Status == 1 {
		return admin, errcode.AccountIsLocked
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return admin, err
	}
	return admin, nil
}
