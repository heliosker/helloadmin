package models

import "gorm.io/gorm"

type ConfigModule struct {
	Model
	Module string `json:"module"`
}

func (c ConfigModule) TableName() string {
	return "hi_config_module"
}

func (c ConfigModule) List(db *gorm.DB) ([]*ConfigModule, error) {
	var modules []*ConfigModule
	if err := db.Find(&modules).Error; err != nil {
		return nil, err
	}
	return modules, nil
}

func (c ConfigModule) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}

func (c ConfigModule) Update(db *gorm.DB) error {
	db = db.Model(&ConfigModule{}).Where("id = ?", c.ID)
	return db.Updates(c).Error
}

func (c ConfigModule) Delete(db *gorm.DB) error {
	return db.Where("id = ?", c.ID).Delete(&c).Error
}
