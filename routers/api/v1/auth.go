package v1

import (
	"github.com/gin-gonic/gin"
	"helloadmin/models"
	e "helloadmin/pkg/error"
	"helloadmin/pkg/utils"
	"net/http"
)

type Profile struct {
	Username string `json:"username"`
}

func AuthLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if !models.AdminUserExist(email, password) {
		c.JSON(utils.Error(http.StatusUnauthorized, e.ERROR_PASSWORD_FAIL))
		return
	}

	token, err := utils.GetToken(email)
	if err != nil {
	}
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, map[string]string{"token": token}, nil))
}

func AuthLogout(c *gin.Context) {
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, nil, nil))
}
