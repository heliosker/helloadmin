package v1

import (
	"github.com/gin-gonic/gin"
	e "helloadmin/pkg/error"
	"helloadmin/pkg/utils"
	"net/http"
)

func RoleIndex(c *gin.Context) {
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, nil, nil))
}

func RoleStore(c *gin.Context) {
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, nil, nil))
}

func RoleShow(c *gin.Context) {
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, nil, nil))
}

func RoleUpdate(c *gin.Context) {
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, nil, nil))
}

func RoleDestroy(c *gin.Context) {
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, nil, nil))
}
