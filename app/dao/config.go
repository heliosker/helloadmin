package dao

import (
	"helloadmin/app/models"
)

func (d Dao) List(moduleId uint) ([]models.Config, error) {
	cfg := models.Config{ModuleId: moduleId}
	return cfg.List(d.engine)
}

func (d Dao) CreateConfig(moduleId uint, title, key, value, tip string) error {
	cfg := models.Config{
		Title:    title,
		Key:      key,
		Value:    value,
		Tip:      tip,
		ModuleId: moduleId,
	}
	return cfg.Create(d.engine)
}

func (d Dao) UpdateConfig(id, moduleId uint, title, key, value, tip string) error {
	cfg := models.Config{
		Title:    title,
		Key:      key,
		Value:    value,
		Tip:      tip,
		ModuleId: moduleId,
	}
	cfg.ID = id
	return cfg.Update(d.engine)
}

func (d Dao) GetConfigValue(key string) string {
	var cfg = models.Config{Key: key}
	return cfg.Val(d.engine)
}

func (d Dao) DeleteConfig(id uint) error {
	cfg := models.Config{}
	cfg.ID = id
	return cfg.Delete(d.engine)
}
