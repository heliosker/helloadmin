package service

import (
	"helloadmin/app/models"
)

type ConfigReq struct {
	Module string `form:"module" binding:"max=50"`
}

func (svc *Service) GetConfigByGroup(req ConfigReq) ([]models.ConfigRet, error) {
	return svc.dao.GetConfig(req.Module)
}

func (svc *Service) StoreMultiConfig(req models.ConfigStore) error {
	return svc.dao.StoreConfig(req)
}

func (svc *Service) GetValByKey(key string) string {
	if value, err := svc.dao.GetValue(key); err != nil {
		panic(err)
	} else {
		return value
	}
}
