package v1

import (
	"github.com/gin-gonic/gin"
	"helloadmin/app/service"
	"helloadmin/pkg/app"
	"helloadmin/pkg/errcode"
)

type Profile struct {
	Username string `json:"username"`
}

func AuthLogout(c *gin.Context) {
	app.NewResponse(c).Success(nil, app.NoMeta)
}

func AuthLogin(c *gin.Context) {
	req := service.AuthReq{}
	rsp := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &req)
	if !valid {
		rsp.Error(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	if err := svc.CheckAuth(&req); err != nil {
		rsp.Error(errcode.UnauthorizedAuthNotExist.WithDetails(err.Error()))
		return
	}
	token, err := app.GenerateToken(req.Username, req.Password)
	if err != nil {
		rsp.Error(errcode.UnauthorizedTokenGenerate.WithDetails(err.Error()))
	}
	rsp.Success(gin.H{"token": token}, app.NoMeta)
}
