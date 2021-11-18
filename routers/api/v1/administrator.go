package v1

import (
	"github.com/gin-gonic/gin"
	"helloadmin/app/models"
	"helloadmin/app/service"
	"helloadmin/pkg/app"
	"helloadmin/pkg/errcode"
)

func AdminMe(c *gin.Context) {
	rsp := app.NewResponse(c)
	var admin models.AdminUser
	if err := models.DB.Where("username=?", c.Get("username")).Find(&admin).Error; err != nil {
		rsp.Error(errcode.SelectedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(admin, app.NoMeta)
}

func AdminIndex(c *gin.Context) {
	params := service.GetAdminReq{}
	svc := service.New(c.Request.Context())
	// 验证 Todo
	count, _ := svc.CountAdministrators(params)
	users, e := svc.GetAdministrators(params, app.Meta{Count: count})
	if e != nil {
		app.NewResponse(c).Error(errcode.SelectedFail.WithDetails(e.Error()))
		return
	}
	app.NewResponse(c).Success(users, count)
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
		app.NewResponse(c).Error(errcode.CreatedFail.WithDetails(err.Error()))
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
