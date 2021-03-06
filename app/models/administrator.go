package models

import (
	"gorm.io/gorm"
)

type AdminUser struct {
	Model
	Email    string `json:"email"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
	Password string `json:"password"`
	RoleId   int64  `json:"role_id"`
}

func (AdminUser) TableName() string {
	return "hi_admin_users"
}

func (au AdminUser) Pagination(db *gorm.DB, offset, size int) ([]*AdminUser, error) {
	var admin []*AdminUser
	if au.Username != "" {
		db.Where("username = ?", au.Username)
	}
	if au.Email != "" {
		db.Where("email = ?", au.Email)
	}
	if au.Status != -1 {
		db.Where("status = ?", au.Status)
	}
	if err := db.Find(&admin).Offset(offset).Limit(size).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (au AdminUser) Count(db *gorm.DB) (int64, error) {
	var count int64
	if au.Username != "" {
		db.Where("username = ?", au.Username)
	}
	if au.Email != "" {
		db.Where("email = ?", au.Email)
	}
	if au.Status != -1 {
		db.Where("status = ?", au.Status)
	}
	if err := db.Model(&au).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}
