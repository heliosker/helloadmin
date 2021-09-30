package v1

import (
	"github.com/gin-gonic/gin"
	"helloadmin/models"
	e "helloadmin/pkg/error"
	"helloadmin/pkg/utils"
	"net/http"
	"strconv"
)

func VersionIndex(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	var count int64
	var version []models.Version
	models.DB.Model(&version).Count(&count)
	ret := models.DB.Scopes(utils.Paginate(page, size)).Find(&version)
	if ret.Error != nil {
		c.JSON(utils.Error(http.StatusOK, e.ERROR_CREATED_FAIL))
	}
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, version, &utils.Meta{Page: page, Size: size, Total: count}))
}

func VersionStore(c *gin.Context) {
	var version models.Version
	_ = c.ShouldBindJSON(&version)

	err := models.DB.Create(&version).Error
	if err != nil {
		c.JSON(utils.Error(http.StatusOK, e.ERROR_CREATED_FAIL))
		return
	}
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, map[string]string{}, nil))
}
