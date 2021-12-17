package models

import "gorm.io/gorm"

type Role struct {
	Model
	Name     string `json:"name"`
	Describe string `json:"describe"`
}

func (Role) TableName() string {
	return "hi_roles"
}

func (r Role) One(db *gorm.DB) Role {
	var role Role
	db.Where("id = ?", r.ID).Find(&role)
	return role
}

func (r Role) Create(db *gorm.DB) error {
	return db.Create(&r).Error
}

func (r Role) Update(db *gorm.DB) error {
	return db.Model(&Role{}).Where("id = ?", r.ID).Updates(r).Error
}

func (r Role) Delete(db *gorm.DB) error {
	return db.Where("id = ?", r.ID).Delete(&r).Error
}
