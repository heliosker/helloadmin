package v1

import (
	"helloadmin/app/models"
	"helloadmin/app/service"
	"helloadmin/pkg/app"
	"helloadmin/pkg/convert"
	"helloadmin/pkg/errcode"
	"helloadmin/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Role struct {
}

func NewRole() Role {
	return Role{}
}

// 绝色列表
func (r Role) Index(c *gin.Context) {
	p, _ := strconv.Atoi(c.Query("page"))
	s, _ := strconv.Atoi(c.Query("size"))
	var count int64
	var roles []models.Role
	if c.Query("options") != "" {
		models.DB.Find(&roles)
		result := make([]map[string]interface{}, 0, len(roles))
		for i := 0; i < len(roles); i++ {
			ops := map[string]interface{}{"key": roles[i].ID, "value": roles[i].Name}
			result = append(result, ops)
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

// 创建角色
func (r Role) Store(c *gin.Context) {
	var role service.CreateRole
	rsp := app.NewResponse(c)
	if valid, errors := app.BindAndValid(c, &role); !valid {
		rsp.Error(errcode.InvalidParams.WithDetails(errors.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.CreateRole(role); err != nil {
		rsp.Error(errcode.CreatedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(nil, app.NoMeta)
}

// 查询角色
func (r Role) Show(c *gin.Context) {
	roleId := convert.StrTo(c.Param("id")).MustUInt()
	svc := service.New(c.Request.Context())
	rsp := app.NewResponse(c)
	role := svc.FindRole(roleId)
	rsp.Success(role, app.NoMeta)
}

// 角色修改
func (r Role) Update(c *gin.Context) {
	roleId := convert.StrTo(c.Param("id")).MustUInt()
	param := service.UpdateRole{}
	rsp := app.NewResponse(c)
	if valid, errors := app.BindAndValid(c, &param); !valid {
		rsp.Error(errcode.InvalidParams.WithDetails(errors.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.UpdateRole(roleId, param); err != nil {
		rsp.Error(errcode.UpdatedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(nil, app.NoMeta)
}

// 角色删除
func (r Role) Destroy(c *gin.Context) {
	roleId := convert.StrTo(c.Param("id")).MustUInt()
	svc := service.New(c.Request.Context())
	rsp := app.NewResponse(c)
	if err := svc.DeleteRole(roleId); err != nil {
		rsp.Error(errcode.DeletedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(nil, app.NoMeta)
}

// 新增角色菜单
func (r Role) CreatedMenu(c *gin.Context) {
	roleId := convert.StrTo(c.Param("id")).MustUInt()
	menu := service.CreateRoleMenu{}

	rsp := app.NewResponse(c)
	if valid, errors := app.BindAndValid(c, &menu); !valid {
		rsp.Error(errcode.InvalidParams.WithDetails(errors.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.CreateRoleMenu(roleId, menu); err != nil {
		rsp.Error(errcode.CreatedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(nil, app.NoMeta)
}

// 修改角色菜单
func (r Role) UpdatedMenu(c *gin.Context) {
	menu := service.CreateRoleMenu{}
	roleId := convert.StrTo(c.Param("id")).MustUInt()
	// Valid
	rsp := app.NewResponse(c)
	if valid, errors := app.BindAndValid(c, &menu); !valid {
		rsp.Error(errcode.InvalidParams.WithDetails(errors.Error()))
		return
	}
	// Service
	svc := service.New(c.Request.Context())
	if err := svc.DeleteRoleMenu(roleId); err != nil {
		rsp.Error(errcode.DeletedFail.WithDetails(err.Error()))
		return
	}
	if err := svc.CreateRoleMenu(roleId, menu); err != nil {
		rsp.Error(errcode.CreatedFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(nil, app.NoMeta)
}
