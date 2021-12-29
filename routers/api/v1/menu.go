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

// Index 菜单列表 godoc
// @Summary 菜单列表
// @Param options query int 0 "int enums" Enums(0, 1)
// @Tags 菜单
// @Accept json
// @Produce json
// @Success 200 {object} []service.MenuTreeMap
// @Failure 500 {object} errcode.Error
// @Router /menus [get]
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

// Show 查看菜单 godoc
// @Summary 查询菜单
// @Param id path int true "Menu ID"
// @Tags 菜单
// @Accept json
// @Produce json
// @Success 200 {object} models.Menu
// @Failure 500 {object} errcode.Error
// @Router /menus/{id} [get]
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

// Store 创建菜单 godoc
// @Summary 创建菜单
// @Param parent_id body int false "menu ID"
// @Param label body string false "string valid"
// @Param Name body string true "string valid"
// @Param Icon body string true "string valid"
// @Param Path body string true "string valid"
// @Param Redirect body string false "string valid"
// @Param sort body int false "string default" default(50)
// @Param show body int false "string enums" Enums(0,1)
// @Tags 菜单
// @Accept json
// @Produce json
// @Success 200 {object} models.Menu
// @Failure 500 {object} errcode.Error
// @Router /menus [post]
func (m Menu) Store(c *gin.Context) {
	var menu models.Menu
	_ = c.ShouldBindJSON(&menu)
	if err := models.DB.Create(&menu).Error; err != nil {
		app.NewResponse(c).Error(errcode.CreatedFail.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).Success(menu, app.NoMeta)
}

// Update 修改菜单 godoc
// @Summary 修改菜单
// @Tags 菜单
// @Accept json
// @Produce json
// @Success 200 {object} nil
// @Failure 500 {object} errcode.Error
// @Router /menus/{id} [put]
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

// Delete 删除菜单 godoc
// @Summary 删除菜单
// @Tags 菜单
// @Accept json
// @Produce json
// @Success 200 {object} nil
// @Failure 500 {object} errcode.Error
// @Router /menus/{id} [delete]
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
