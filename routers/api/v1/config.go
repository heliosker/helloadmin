package v1

import (
	"helloadmin/app/service"
	"helloadmin/pkg/app"
	"helloadmin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Config struct {
}

func NewConfig() Config {
	return Config{}
}

func (f Config) Index(c *gin.Context) {
	req := service.ConfigListReq{}
	rsp := app.NewResponse(c)
	if valid, errors := app.BindAndValid(c, &req); !valid {
		rsp.Error(errcode.InvalidParams.WithDetails(errors.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	config, err := svc.GetConfig(req)
	if err != nil {
		rsp.Error(errcode.SelectedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(config, app.NoMeta)
}

func (f Config) Create(c *gin.Context) {
	param := service.CreateConfig{}
	rsp := app.NewResponse(c)
	if valid, errors := app.BindAndValid(c, &param); !valid {
		rsp.Error(errcode.InvalidParams.WithDetails(errors.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.CreateConfig(param); err != nil {
		rsp.Error(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	rsp.Success(nil, app.NoMeta)
}

func (f Config) Update(c *gin.Context) {
	param := service.UpdateMultiConfig{}
	rsp := app.NewResponse(c)
	if valid, errors := app.BindAndValid(c, &param); !valid {
		rsp.Error(errcode.InvalidParams.WithDetails(errors.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.UpdateMultiConfig(param); err != nil {
		rsp.Error(errcode.UpdatedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(nil, app.NoMeta)
}

func (f Config) Delete(c *gin.Context) {
	param := service.DeleteConfig{}
	rsp := app.NewResponse(c)
	if valid, errors := app.BindAndValid(c, &param); !valid {
		rsp.Error(errcode.InvalidParams.WithDetails(errors.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.DeleteConfig(param); err != nil {
		rsp.Error(errcode.DeletedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(nil, app.NoMeta)
}
