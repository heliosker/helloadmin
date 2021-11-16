package v1

import (
	"github.com/gin-gonic/gin"
	"helloadmin/app/models"
	"helloadmin/pkg/app"
	"helloadmin/pkg/errcode"
	"helloadmin/pkg/utils"
)

type Profile struct {
	Username string `json:"username"`
}

func AuthLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if err := models.AdminUserLogin(email, password); err != nil {
		app.NewResponse(c).Error(err)
		return
	}
	token, err := utils.GetToken(email)
	if err != nil {
		app.NewResponse(c).Error(errcode.GenerateTokenError)
		return
	}
	app.NewResponse(c).Success(gin.H{"token": token}, app.NoMeta)
}

func AuthLogout(c *gin.Context) {
	app.NewResponse(c).Success(nil, app.NoMeta)
}
