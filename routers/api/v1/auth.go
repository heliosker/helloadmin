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

// AuthLogout 退出
// @Summary 管理员退出
// @Param Authorization header string true "{access_token}"
// @Tags 登录与退出
// @Accept json
// @Produce json
// @Success 200 {object} nil
// @Failure 401 {object} nil
// @Router /auth/logout [delete]
func AuthLogout(c *gin.Context) {
	app.NewResponse(c).Success(nil, app.NoMeta)
}

// AuthLogin 登录 godoc
// @Summary 管理员登录
// @Param username formData string true "string default" default(admin)
// @Param password formData string true "string valid" minlength(5)  maxlength(10)
// @Tags 登录与退出
// @Accept json
// @Produce json
// @Success 200 {object} nil
// @Failure 401 {object} errcode.Error
// @Router /auth/login [post]
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
	rsp.Success(gin.H{"access_token": token}, app.NoMeta)
}
