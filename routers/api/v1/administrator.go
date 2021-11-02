package v1

import (
	"github.com/gin-gonic/gin"
	"helloadmin/models"
	e "helloadmin/pkg/error"
	"helloadmin/pkg/utils"
	"net/http"
	"strconv"
)

func AdminIndex(c *gin.Context) {
	p, _ := strconv.Atoi(c.Query("page"))
	s, _ := strconv.Atoi(c.Query("size"))
	var count int64
	var admin []models.AdminUser
	models.DB.Model(&admin).Count(&count)

	ret := models.DB.Scopes(utils.Paginate(p, s)).Find(&admin)
	if ret.Error != nil {
		c.JSON(utils.Error(http.StatusOK, e.ERROR_SELECT_FAIL))
		return
	}
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, admin, &utils.Meta{Page: p, Size: s, Total: count}))
}

func AdminShow(c *gin.Context) {
	id := c.Param("id")
	var admin models.AdminUser
	if ret := models.DB.Where("id", id).Find(&admin); ret.Error != nil {
		c.JSON(utils.Error(http.StatusOK, e.ERROR_SELECT_FAIL))
		return
	}
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, admin, nil))
}

func AdminStore(c *gin.Context) {
	var admin models.AdminUser
	_ = c.ShouldBindJSON(&admin)
	if err := models.DB.Create(&admin).Error; err != nil {
		c.JSON(utils.Error(http.StatusOK, e.ERROR_CREATED_FAIL))
		return
	}
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, admin, nil))
}

func AdminUpdate(c *gin.Context) {
	id := c.Param("id")
	var admin models.AdminUser
	_ = c.ShouldBindJSON(&admin)
	if ret := models.DB.Where("id", id).Updates(admin); ret.Error != nil {
		c.JSON(utils.Error(http.StatusOK, e.ERROR_UPDATED_FAIL))
		return
	}
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, admin, nil))
}

func AdminDelete(c *gin.Context) {
	id := c.Param("id")
	var admin models.AdminUser
	if ret := models.DB.Unscoped().Where("id", id).Delete(&admin); ret.Error != nil {
		c.JSON(utils.Error(http.StatusOK, e.ERROR_DELETED_FAIL))
	}
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, nil, nil))
}
