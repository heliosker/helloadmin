package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helloadmin/api"
	"helloadmin/internal/ecode"
	"helloadmin/internal/service"
	"net/http"
	"strconv"
)

type RoleHandler struct {
	*Handler
	roleService service.RoleService
}

func NewRoleHandler(handler *Handler, roleService service.RoleService) *RoleHandler {
	return &RoleHandler{
		Handler:     handler,
		roleService: roleService,
	}
}

// StoreRole godoc
// @Summary 创建角色
// @Schemes
// @Description 创建角色
// @Tags 角色模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body api.RoleCreateRequest true "params"
// @Success 200 {object} api.Response
// @Router /role [post]
func (r *RoleHandler) StoreRole(ctx *gin.Context) {
	req := new(api.RoleCreateRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		r.logger.WithContext(ctx).Error("RoleHandler.StoreRole ShouldBindJSON error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if err := r.roleService.CreateRole(ctx, req); err != nil {
		r.logger.WithContext(ctx).Error("roleService.CreateRole error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, nil)
}

// GetRole godoc
// @Summary 角色列表
// @Schemes
// @Description 查询角色列表
// @Tags 角色模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request query api.RoleFindRequest true "params"
// @Success 200 {object} api.Response
// @Router /role [get]
func (r *RoleHandler) GetRole(ctx *gin.Context) {
	req := new(api.RoleFindRequest)
	if err := ctx.ShouldBindQuery(req); err != nil {
		r.logger.WithContext(ctx).Error("RoleHandler.GetRole error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	role, err := r.roleService.SearchRole(ctx, req)
	if err != nil {
		r.logger.WithContext(ctx).Error("roleService.SearchRole error", zap.Error(err))
		return
	}
	api.Success(ctx, role)
}

// ShowRole godoc
// @Summary 查询角色
// @Schemes
// @Description 查询单个角色信息
// @Tags 角色模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "角色ID"
// @Success 200 {object} api.Response
// @Router /role/{id} [get]
func (r *RoleHandler) ShowRole(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if role, err := r.roleService.GetRoleById(ctx, id); err != nil {
		r.logger.WithContext(ctx).Error("roleService.GetRoleById error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	} else {
		api.Success(ctx, role)
	}
}

// UpdateRole godoc
// @Summary 修改角色
// @Schemes
// @Description 修改单个角色信息
// @Tags 角色模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "角色ID"
// @Param request body api.RoleUpdateRequest true "params"
// @Success 200 {object} api.Response
// @Router /role/{id} [put]
func (r *RoleHandler) UpdateRole(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	req := new(api.RoleUpdateRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		r.logger.WithContext(ctx).Error("RoleHandler.ShowRole error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, ecode.ErrBadRequest)
		return
	}
	if err := r.roleService.UpdateRole(ctx, id, req); err != nil {
		r.logger.WithContext(ctx).Error("roleService.UpdateRole error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, nil)
}

// UpdateRoleMenu godoc
// @Summary 修改角色权限
// @Schemes
// @Description 修改单个角色权限
// @Tags 角色模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "角色ID"
// @Param request body api.RoleMenuRequest true "params"
// @Success 200 {object} api.Response
// @Router /role/{id}/menu [put]
func (r *RoleHandler) UpdateRoleMenu(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	req := new(api.RoleMenuRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		r.logger.WithContext(ctx).Error("RoleHandler.ShowRole error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if err := r.roleService.UpdateRoleMenu(ctx, id, req); err != nil {
		r.logger.WithContext(ctx).Error("roleService.UpdateRole error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, nil)
}

// DeleteRole godoc
// @Summary 删除角色
// @Schemes
// @Description 删除单个角色
// @Tags 角色模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "角色ID"
// @Success 200 {object} api.Response
// @Router /role/{id} [delete]
func (r *RoleHandler) DeleteRole(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := r.roleService.DeleteRole(ctx, id); err != nil {
		r.logger.WithContext(ctx).Error("roleService.DeleteRole error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, nil)
}
