package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"helloadmin/models"
	"helloadmin/pkg/app"
	"helloadmin/pkg/errcode"
)

type Menu struct {
}

func NewMenu() Menu {
	return Menu{}
}

func (m Menu) Index(c *gin.Context) {

	app.NewResponse(c).Success(nil, 5)
}

func (m Menu) Show(c *gin.Context) {
	id := c.Param("id")
	var menu models.Menu
	rsp := app.NewResponse(c)
	models.DB.Where("id", id).Find(&menu)
	if menu.ID == 0 {
		rsp.Error(errcode.NotFound)
		return
	}
	rsp.Success(menu, app.NoMeta)
}

func (m Menu) Store(c *gin.Context) {
	var menu models.Menu
	_ = c.ShouldBindJSON(&menu)
	fmt.Println(menu)
	if err := models.DB.Create(&menu).Error; err != nil {
		app.NewResponse(c).Error(errcode.CreatedFail.WithDetails(err.Error()))
	}
	app.NewResponse(c).Success(menu, app.NoMeta)
}

func (m Menu) Update(c *gin.Context) {
	id := c.Param("id")
	var menu models.Menu
	_ = c.ShouldBindJSON(&menu)
	rsp := app.NewResponse(c)
	if err := models.DB.Where("id", id).Updates(&menu).Error; err != nil {
		rsp.Error(errcode.UpdatedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(nil, app.NoMeta)
}

func (m Menu) Delete(c *gin.Context) {
	id := c.Param("id")
	var menu models.Menu
	rsp := app.NewResponse(c)
	if err := models.DB.Debug().Where("id", id).Delete(&menu).Error; err != nil {
		rsp.Error(errcode.DeletedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(nil, app.NoMeta)
}
