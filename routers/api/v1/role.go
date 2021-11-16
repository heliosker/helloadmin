package v1

import (
	"github.com/gin-gonic/gin"
	"helloadmin/app/models"
	"helloadmin/pkg/app"
	"helloadmin/pkg/errcode"
	"helloadmin/pkg/utils"
	"strconv"
)

type Role struct {
}

func NewRole() Role {
	return Role{}
}

func (r Role) Index(c *gin.Context) {
	p, _ := strconv.Atoi(c.Query("page"))
	s, _ := strconv.Atoi(c.Query("size"))
	var count int64
	var roles []models.Role
	if c.Query("options") != "" {
		models.DB.Find(&roles)
		var result = map[uint]string{}
		for i := 0; i < len(roles); i++ {
			result[roles[i].ID] = roles[i].Name
		}
		app.NewResponse(c).Success(result, app.NoMeta)
		return
	}

	models.DB.Model(&roles).Count(&count)
	ret := models.DB.Scopes(utils.Paginate(p, s)).Find(&roles)
	if ret.Error != nil {
		app.NewResponse(c).Error(errcode.SelectedFail.WithDetails(ret.Error.Error()))
		return
	}
	app.NewResponse(c).Success(roles, count)
}

func (r Role) Store(c *gin.Context) {
	var role models.Role
	_ = c.ShouldBindJSON(&role)
	err := models.DB.Create(&role).Error
	rsp := app.NewResponse(c)
	if err != nil {
		rsp.Error(errcode.CreatedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(role, app.NoMeta)
}

func (r Role) Show(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	_ = models.DB.Where("id", id).Find(&role)
	rsp := app.NewResponse(c)
	if role.ID == 0 {
		rsp.Error(errcode.NotFound)
		return
	}
	rsp.Success(role, app.NoMeta)
}

func (r Role) Update(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	_ = c.ShouldBindJSON(&role)
	ret := models.DB.Where("id", id).Updates(role)
	rsp := app.NewResponse(c)
	if ret.Error != nil {
		rsp.Error(errcode.UpdatedFail.WithDetails(ret.Error.Error()))
		return
	}
	rsp.Success(nil, app.NoMeta)
}

func (r Role) Destroy(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	ret := models.DB.Where("id", id).Delete(&role)
	rsp := app.NewResponse(c)
	if ret.Error != nil {
		rsp.Error(errcode.DeletedFail.WithDetails(ret.Error.Error()))
		return
	}
	rsp.Success(nil, app.NoMeta)
}
