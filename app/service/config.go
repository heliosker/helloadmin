package service

import (
	"encoding/json"
	"helloadmin/app/models"
)

type ConfigListReq struct {
	ModuleId uint `form:"module_id" binding:"max=50"`
}

type CreateConfig struct {
	ModuleId uint   `form:"module_id" binding:"required"`
	Title    string `form:"title" binding:"required,max=30"`
	Key      string `form:"key" binding:"required,max=60"`
	Value    string `form:"value" binding:"required,max=255"`
	Tip      string `form:"tip" binding:"required,max=255"`
}

type UpdateConfig struct {
	ID       uint   `form:"id" binding:"required"`
	ModuleId uint   `form:"module_id"`
	Title    string `form:"title" binding:"max=30"`
	Key      string `form:"key" binding:"max=60"`
	Value    string `form:"value" binding:max=255"`
	Tip      string `form:"tip" binding:"max=255"`
}

type UpdateMultiConfig struct {
	UpdateConfig []UpdateConfig
}

func (i *UpdateMultiConfig) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &i.UpdateConfig)
}

type DeleteConfig struct {
	ID uint `form:"id" binding:"required"`
}

func (svc *Service) GetConfig(r ConfigListReq) ([]models.Config, error) {
	return svc.dao.List(r.ModuleId)
}

func (svc *Service) UpdateMultiConfig(items UpdateMultiConfig) error {
	for _, item := range items.UpdateConfig {
		if err := svc.dao.UpdateConfig(item.ID, item.ModuleId, item.Title, item.Key, item.Value, item.Tip); err != nil {
			return err
		}
	}
	return nil
}

func (svc *Service) CreateConfig(r CreateConfig) error {
	return svc.dao.CreateConfig(r.ModuleId, r.Title, r.Key, r.Value, r.Tip)
}

func (svc *Service) GetConfigValue(key string) string {
	return svc.dao.GetConfigValue(key)
}

func (svc *Service) DeleteConfig(r DeleteConfig) error {
	return svc.dao.DeleteConfig(r.ID)
}
