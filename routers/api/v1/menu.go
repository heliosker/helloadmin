package v1

import (
	"github.com/gin-gonic/gin"
	"helloadmin/app/models"
	"helloadmin/app/service"
	"helloadmin/pkg/app"
	"helloadmin/pkg/errcode"
)

type Menu struct {
}

func NewMenu() Menu {
	return Menu{}
}

func (m Menu) Index(c *gin.Context) {
	params := service.MenuListReq{}
	// 验证 Todo
	svc := service.New(c.Request.Context())
	rsp := app.NewResponse(c)
	if c.Query("options") != "" {
		option, e := svc.GetOptions()
		if e != nil {
			rsp.Error(errcode.SelectedFail.WithDetails(e.Error()))
			return
		}
		rsp.Success(option, app.NoMeta)
	} else {
		//boo, errors := app.BindAndValid(c, &params)
		if menus, e := svc.GetTreeMenu(&params); e != nil {
			rsp.Error(errcode.SelectedFail.WithDetails(e.Error()))
		} else {
			rsp.Success(menus, app.NoMeta)
		}
	}
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
	if err := models.DB.Create(&menu).Error; err != nil {
		app.NewResponse(c).Error(errcode.CreatedFail.WithDetails(err.Error()))
		return
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
	if err := models.DB.Where("id", id).Delete(&menu).Error; err != nil {
		rsp.Error(errcode.DeletedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(nil, app.NoMeta)
}
