package service

import (
	"helloadmin/app/models"
	"helloadmin/pkg/app"
)

type GetAdminReq struct {
	UserName string `form:"username" binding:"max=100"`
	Email    string `form:"email" binding:"max=100"`
	Status   int    `form:"status"`
}

func (svc *Service) GetAdministrators(req GetAdminReq, meta app.Meta) ([]*models.AdminUser, error) {
	return svc.dao.GetAdministrator(req.Email, req.UserName, req.Status, meta.Page, meta.Size)
}

func (svc *Service) CountAdministrators(req GetAdminReq) (int64, error) {
	return svc.dao.CountAdministrator(req.Email, req.UserName, req.Status)
}
