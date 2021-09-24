package api

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

	if !models.CheckAdminUser(username, password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    e.ERROR_PASSWORD_FAIL,
			"message": e.Message(e.ERROR_PASSWORD_FAIL),
		})
		return
	}

	token, err := utils.GetToken(username)
	if err != nil {
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": e.Message(e.SUCCESS),
		"data":    token,
	})
}
