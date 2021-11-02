package v1

import (
	"github.com/gin-gonic/gin"
	"helloadmin/models"
	e "helloadmin/pkg/error"
	"helloadmin/pkg/utils"
	"net/http"
	"strconv"
)

func RoleIndex(c *gin.Context) {
	p, _ := strconv.Atoi(c.Query("page"))
	s, _ := strconv.Atoi(c.Query("size"))
	var count int64
	var roles []models.Role

	models.DB.Model(&roles).Count(&count)
	ret := models.DB.Scopes(utils.Paginate(p, s)).Find(&roles)
	if ret.Error != nil {
		c.JSON(utils.Error(http.StatusOK, e.ERROR_SELECT_FAIL))
		return
	}
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, roles, &utils.Meta{Page: p, Size: s, Total: count}))
}

func RoleStore(c *gin.Context) {
	var role models.Role
	_ = c.ShouldBindJSON(&role)
	err := models.DB.Create(&role).Error
	if err != nil {
		c.JSON(utils.Error(http.StatusOK, e.ERROR_CREATED_FAIL))
		return
	}
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, role, nil))
}

func RoleShow(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	ret := models.DB.Where("id", id).Find(&role)
	if ret.Error != nil || role.ID == 0 {
		c.JSON(utils.Error(http.StatusOK, e.ERROR_SELECT_FAIL))
		return
	}
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, role, nil))
}

func RoleUpdate(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	_ = c.ShouldBindJSON(&role)
	ret := models.DB.Where("id", id).Updates(role)
	if ret.Error != nil {
		c.JSON(utils.Error(http.StatusOK, e.ERROR_UPDATED_FAIL))
		return
	}

	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, nil, nil))
}

func RoleDestroy(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	ret := models.DB.Where("id", id).Delete(&role)
	if ret.Error != nil {
		c.JSON(utils.Error(http.StatusOK, e.ERROR_DELETED_FAIL))
		return
	}
	c.JSON(utils.Success(http.StatusOK, e.SUCCESS, nil, nil))
}
