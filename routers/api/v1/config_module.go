package v1

import (
	"helloadmin/app/service"
	"helloadmin/pkg/app"
	"helloadmin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type ConfigModule struct {
}

func NewConfigModule() ConfigModule {
	return ConfigModule{}
}

func (cm ConfigModule) Index(c *gin.Context) {
	rsp := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	ret, err := svc.GetConfigModules()
	if err != nil {
		rsp.Error(errcode.SelectedFail.WithDetails(err.Error()))
	}
	rsp.Success(ret, app.NoMeta)

}

func (cm ConfigModule) Create(c *gin.Context) {
	param := service.CreateConfigModule{}
	rsp := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		rsp.Error(errcode.InvalidParams.WithDetails(errors.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.CreateConfigModule(&param); err != nil {
		rsp.Error(errcode.CreatedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(nil, app.NoMeta)
}

func (cm ConfigModule) Update(c *gin.Context) {
	param := service.UpdateConfigModule{}
	rsp := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		rsp.Error(errcode.InvalidParams.WithDetails(errors.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.UpdateConfigModule(&param); err != nil {
		rsp.Error(errcode.UpdatedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(nil, app.NoMeta)
}

func (cm ConfigModule) Delete(c *gin.Context) {
	param := service.DeleteConfigModule{}
	rsp := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		rsp.Error(errcode.InvalidParams.WithDetails(errors.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.DeleteConfigModule(&param); err != nil {
		rsp.Error(errcode.DeletedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(nil, app.NoMeta)
}
