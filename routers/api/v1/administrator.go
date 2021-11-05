package v1

import (
	"github.com/gin-gonic/gin"
	"helloadmin/models"
	"helloadmin/pkg/app"
	"helloadmin/pkg/errcode"
	"helloadmin/pkg/utils"
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
		app.NewResponse(c).Error(errcode.SelectedFail)
		return
	}
	app.NewResponse(c).Success(admin, count)
}

func AdminShow(c *gin.Context) {
	id := c.Param("id")
	var admin models.AdminUser
	if ret := models.DB.Where("id", id).Find(&admin); ret.Error != nil {
		app.NewResponse(c).Error(errcode.SelectedFail)
		return
	}
	app.NewResponse(c).Success(admin, app.NoMeta)
}

func AdminStore(c *gin.Context) {
	var admin models.AdminUser
	_ = c.ShouldBindJSON(&admin)
	if err := models.DB.Create(&admin).Error; err != nil {
		app.NewResponse(c).Error(errcode.CreatedFail)
		return
	}
	app.NewResponse(c).Success(admin, app.NoMeta)
}

func AdminUpdate(c *gin.Context) {
	id := c.Param("id")
	var admin models.AdminUser
	_ = c.ShouldBindJSON(&admin)
	if ret := models.DB.Where("id", id).Updates(admin); ret.Error != nil {
		app.NewResponse(c).Error(errcode.UpdatedFail)
		return
	}
	app.NewResponse(c).Success(admin, app.NoMeta)
}

func AdminDelete(c *gin.Context) {
	id := c.Param("id")
	var admin models.AdminUser
	if ret := models.DB.Where("id", id).Delete(&admin); ret.Error != nil {
		app.NewResponse(c).Error(errcode.DeletedFail)
		return
	}
	app.NewResponse(c).Success(nil, app.NoMeta)
}
