package models

import (
	"gorm.io/gorm"
)

type Config struct {
	Model
	ModuleId uint   `json:"module_id"`
	Title    string `json:"title"`
	Key      string `json:"key"`
	Value    string `json:"value"`
	Tip      string `json:"tip"`
}

func (Config) TableName() string {
	return "hi_config"
}

func (c Config) List(db *gorm.DB) ([]Config, error) {
	var cfg []Config
	if c.ModuleId != 0 {
		db = db.Where("`module_id` = ?", c.ModuleId)
	}
	if err := db.Find(&cfg).Error; err != nil {
		return cfg, err
	}
	return cfg, nil
}

func (c Config) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}

func (c Config) Update(db *gorm.DB) error {
	return db.Model(&Config{}).Where("id = ?", c.ID).Updates(c).Error
}

func (c Config) Delete(db *gorm.DB) error {
	return db.Where("id = ?", c.ID).Delete(&c).Error
}

func (c Config) Val(db *gorm.DB) string {
	var cfg Config
	db.Where("`key` = ?", c.Key).Find(&cfg)
	return cfg.Value
}
