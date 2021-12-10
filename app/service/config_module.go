package service

import "helloadmin/app/models"

type CreateConfigModule struct {
	Module string `form:"module" binding:"required,min=2,max=10"`
}

type UpdateConfigModule struct {
	ID     uint   `form:"id" binding:"required,gte=1"`
	Module string `form:"module" binding:"required,max=10"`
}

type DeleteConfigModule struct {
	ID uint `form:"id" binding:"required,gte=1"`
}

func (svc *Service) GetConfigModules() ([]*models.ConfigModule, error) {
	return svc.dao.GetConfigModules()
}

func (svc *Service) CreateConfigModule(req *CreateConfigModule) error {
	return svc.dao.CreateConfigModule(req.Module)
}

func (svc *Service) UpdateConfigModule(req *UpdateConfigModule) error {
	return svc.dao.UpdateConfigModule(req.ID, req.Module)
}

func (svc *Service) DeleteConfigModule(req *DeleteConfigModule) error {
	return svc.dao.DeleteConfigModule(req.ID)
}
