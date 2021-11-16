package service

import (
	"context"
	"helloadmin/app/dao"
	"helloadmin/app/models"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(models.DB)
	return svc
}
