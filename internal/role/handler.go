package role

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helloadmin/internal/api"
	"helloadmin/internal/ecode"
	"helloadmin/pkg/log"
	"net/http"
	"strconv"
)

type Handler struct {
	logger *log.Logger
	svc    Service
}

func NewHandler(log *log.Logger, svc Service) *Handler {
	return &Handler{
		logger: log,
		svc:    svc,
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
// @Param request body RoleCreateRequest true "params"
// @Success 200 {object} api.Response
// @Router /role [post]
func (r *Handler) StoreRole(ctx *gin.Context) {
	req := new(RoleCreateRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		r.logger.WithContext(ctx).Error("RoleHandler.StoreRole ShouldBindJSON error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if err := r.svc.CreateRole(ctx, req); err != nil {
		r.logger.WithContext(ctx).Error("svc.CreateRole error", zap.Error(err))
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
// @Param request query RoleFindRequest true "params"
// @Success 200 {object} api.Response
// @Router /role [get]
func (r *Handler) GetRole(ctx *gin.Context) {
	req := new(RoleFindRequest)
	if err := ctx.ShouldBindQuery(req); err != nil {
		r.logger.WithContext(ctx).Error("RoleHandler.GetRole error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	role, err := r.svc.SearchRole(ctx, req)
	if err != nil {
		r.logger.WithContext(ctx).Error("svc.SearchRole error", zap.Error(err))
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
func (r *Handler) ShowRole(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if role, err := r.svc.GetRoleById(ctx, id); err != nil {
		r.logger.WithContext(ctx).Error("svc.GetRoleById error", zap.Error(err))
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
// @Param request body RoleUpdateRequest true "params"
// @Success 200 {object} api.Response
// @Router /role/{id} [put]
func (r *Handler) UpdateRole(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	req := new(RoleUpdateRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		r.logger.WithContext(ctx).Error("RoleHandler.ShowRole error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, ecode.ErrBadRequest)
		return
	}
	if err := r.svc.UpdateRole(ctx, id, req); err != nil {
		r.logger.WithContext(ctx).Error("svc.UpdateRole error", zap.Error(err))
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
// @Param request body RoleMenuRequest true "params"
// @Success 200 {object} api.Response
// @Router /role/{id}/menu [put]
func (r *Handler) UpdateRoleMenu(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	req := new(RoleMenuRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		r.logger.WithContext(ctx).Error("RoleHandler.ShowRole error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if err := r.svc.UpdateRoleMenu(ctx, id, req); err != nil {
		r.logger.WithContext(ctx).Error("svc.UpdateRole error", zap.Error(err))
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
func (r *Handler) DeleteRole(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := r.svc.DeleteRole(ctx, id); err != nil {
		r.logger.WithContext(ctx).Error("svc.DeleteRole error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, nil)
}
