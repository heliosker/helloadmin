package menu

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helloadmin/internal/api"
	"helloadmin/pkg/log"
	"net/http"
	"strconv"
)

type Handler struct {
	log *log.Logger
	svc Service
}

func NewHandler(log *log.Logger, svc Service) *Handler {
	return &Handler{
		log: log,
		svc: svc,
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
// @Param request body CreateRequest true "params"
// @Success 200 {object} api.Response
// @Router /menu [post]
func (m *Handler) StoreMenu(ctx *gin.Context) {
	req := new(CreateRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if err := m.svc.CreateMenu(ctx, req); err != nil {
		m.log.WithContext(ctx).Error("svc.CreateMenu error", zap.Error(err))
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
// @Param request query FindRequest true "params"
// @Success 200 {object} Response
// @Router /menu [get]
func (m *Handler) GetMenu(ctx *gin.Context) {
	req := new(FindRequest)
	if err := ctx.ShouldBindQuery(req); err != nil {
		m.log.WithContext(ctx).Error("MenuHandler.GetMenu error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if resp, err := m.svc.SearchMenu(ctx, req); err != nil {
		m.log.WithContext(ctx).Error("svc.SearchMenu error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
	} else {
		api.Success(ctx, resp)
	}
}

// GetOption godoc
// @Summary 菜单选项
// @Schemes
// @Description 菜单下拉选项
// @Tags 菜单模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request query OptionRequest true "params"
// @Success 200 {object} []Option
// @Router /menu/option [get]
func (m *Handler) GetOption(ctx *gin.Context) {
	req := new(OptionRequest)
	if err := ctx.ShouldBindQuery(req); err != nil {
		m.log.WithContext(ctx).Error("MenuHandler.Options error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if resp, err := m.svc.Options(ctx, req); err != nil {
		m.log.WithContext(ctx).Error("svc.Options error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
	} else {
		fmt.Println(resp)
		fmt.Println("-------")
		api.Success(ctx, resp)
	}
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
func (m *Handler) ShowMenu(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if menu, err := m.svc.GetMenuById(ctx, id); err != nil {
		m.log.WithContext(ctx).Error("svc.GetMenuById error", zap.Error(err))
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
// @Param request body UpdateRequest true "params"
// @Success 200 {object} api.Response
// @Router /menu/{id} [put]
func (m *Handler) UpdateMenu(ctx *gin.Context) {
	req := new(UpdateRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		m.log.WithContext(ctx).Error("MenuHandler.ShowMenu error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := m.svc.UpdateMenu(ctx, id, req); err != nil {
		m.log.WithContext(ctx).Error("svc.UpdateMenu error", zap.Error(err))
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
func (m *Handler) DeleteMenu(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := m.svc.DeleteMenu(ctx, id); err != nil {
		m.log.WithContext(ctx).Error("svc.DeleteMenu error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, nil)
}
