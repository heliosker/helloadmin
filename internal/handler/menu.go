package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helloadmin/api"
	"helloadmin/internal/service"
	"net/http"
	"strconv"
)

type MenuHandler struct {
	*Handler
	menuService service.MenuService
}

func NewMenuHandler(handler *Handler, menuService service.MenuService) *MenuHandler {
	return &MenuHandler{
		Handler:     handler,
		menuService: menuService,
	}
}

// StoreMenu godoc
// @Summary 创建菜单
// @Schemes
// @Description 创建菜单
// @Tags 菜单模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body api.MenuCreateRequest true "params"
// @Success 200 {object} api.Response
// @Router /menu [post]
func (m *MenuHandler) StoreMenu(ctx *gin.Context) {
	req := new(api.MenuCreateRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if err := m.menuService.CreateMenu(ctx, req); err != nil {
		m.logger.WithContext(ctx).Error("menuService.CreateMenu error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, nil)
}

// GetMenu godoc
// @Summary 菜单列表
// @Schemes
// @Description 查询菜单列表
// @Tags 菜单模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request query api.MenuFindRequest true "params"
// @Success 200 {object} api.Response
// @Router /menu [get]
func (m *MenuHandler) GetMenu(ctx *gin.Context) {
	req := new(api.MenuFindRequest)
	if err := ctx.ShouldBindQuery(req); err != nil {
		m.logger.WithContext(ctx).Error("MenuHandler.GetMenu error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	menu, err := m.menuService.SearchMenu(ctx, req)
	if err != nil {
		m.logger.WithContext(ctx).Error("menuService.SearchMenu error", zap.Error(err))
		return
	}
	api.Success(ctx, menu)
}

// ShowMenu godoc
// @Summary 查询菜单
// @Schemes
// @Description 查询单个菜单信息
// @Tags 菜单模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "菜单ID"
// @Success 200 {object} api.Response
// @Router /menu/{id} [get]
func (m *MenuHandler) ShowMenu(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if menu, err := m.menuService.GetMenuById(ctx, id); err != nil {
		m.logger.WithContext(ctx).Error("menuService.GetMenuById error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	} else {
		api.Success(ctx, menu)
	}
}

// UpdateMenu godoc
// @Summary 修改菜单
// @Schemes
// @Description 修改单个菜单信息
// @Tags 菜单模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "菜单ID"
// @Param request body api.MenuUpdateRequest true "params"
// @Success 200 {object} api.Response
// @Router /menu/{id} [put]
func (m *MenuHandler) UpdateMenu(ctx *gin.Context) {
	req := new(api.MenuUpdateRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		m.logger.WithContext(ctx).Error("MenuHandler.ShowMenu error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := m.menuService.UpdateMenu(ctx, id, req); err != nil {
		m.logger.WithContext(ctx).Error("menuService.UpdateMenu error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, nil)
}

// DeleteMenu godoc
// @Summary 删除菜单
// @Schemes
// @Description 删除单个菜单
// @Tags 菜单模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "菜单ID"
// @Success 200 {object} api.Response
// @Router /menu/{id} [delete]
func (m *MenuHandler) DeleteMenu(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := m.menuService.DeleteMenu(ctx, id); err != nil {
		m.logger.WithContext(ctx).Error("menuService.DeleteMenu error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, nil)
}
