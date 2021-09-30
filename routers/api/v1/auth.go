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
	username := c.PostForm("username")
	password := c.PostForm("password")

	if !models.AdminUserExist(username, password) {
		c.JSON(utils.Error(http.StatusUnauthorized, e.ERROR_PASSWORD_FAIL))
		return
	}

	token, err := utils.GetToken(username)
	if err != nil {
	}
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, map[string]string{"token": token}, nil))
}
