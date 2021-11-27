package v1

import (
	"github.com/gin-gonic/gin"
	"helloadmin/app/models"
	"helloadmin/app/service"
	"helloadmin/pkg/app"
	"helloadmin/pkg/errcode"
)

type Config struct {
}

func NewConfig() Config {
	return Config{}
}

func (f Config) Index(c *gin.Context) {
	req := service.ConfigReq{}
	svc := service.New(c.Request.Context())
	rsp := app.NewResponse(c)
	if valid, errors := app.BindAndValid(c, &req); !valid {
		rsp.Error(errcode.InvalidParams.WithDetails(errors.Error()))
		return
	}
	config, err := svc.GetConfigByGroup(req)
	if err != nil {
		rsp.Error(errcode.SelectedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(config, app.NoMeta)
}

func (f Config) Save(c *gin.Context) {
	req := models.ConfigStore{}
	svc := service.New(c.Request.Context())
	rsp := app.NewResponse(c)
	if valid, errors := app.BindAndValid(c, &req); !valid {
		rsp.Error(errcode.InvalidParams.WithDetails(errors.Error()))
		return
	}
	if err := svc.StoreConfig(req); err != nil {
		rsp.Error(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	rsp.Success(gin.H{}, app.NoMeta)
}
