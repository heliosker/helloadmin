package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helloadmin/api"
	"helloadmin/internal/service"
	"net/http"
	"strconv"
)

type DepartmentHandler struct {
	*Handler
	departmentService service.DepartmentService
}

func NewDepartmentHandler(handler *Handler, departmentService service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{
		Handler:           handler,
		departmentService: departmentService,
	}
}

// StoreDepartment godoc
// @Summary 创建部门
// @Schemes
// @Description 创建部门
// @Tags 部门模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body api.DepartmentCreateRequest true "params"
// @Success 200 {object} api.Response
// @Router /department [post]
func (d *DepartmentHandler) StoreDepartment(ctx *gin.Context) {
	req := new(api.DepartmentCreateRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if err := d.departmentService.CreateDepartment(ctx, req); err != nil {
		d.logger.WithContext(ctx).Error("departmentService.CreateDepartment error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, nil)
}

// GetDepartment godoc
// @Summary 部门列表
// @Schemes
// @Description 查询部门列表
// @Tags 部门模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request query api.DepartmentFindRequest true "params"
// @Success 200 {object} api.Response
// @Router /department [get]
func (d *DepartmentHandler) GetDepartment(ctx *gin.Context) {
	req := new(api.DepartmentFindRequest)
	if err := ctx.ShouldBindQuery(req); err != nil {
		d.logger.WithContext(ctx).Error("DepartmentHandler.GetDepartment error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	departments, err := d.departmentService.SearchDepartment(ctx, req)
	if err != nil {
		d.logger.WithContext(ctx).Error("departmentService.SearchDepartment error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, departments)
}

// ShowDepartment godoc
// @Summary 查询部门
// @Schemes
// @Description 查询单个部门信息
// @Tags 部门模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "部门ID"
// @Success 200 {object} api.Response
// @Router /department/{id} [get]
func (d *DepartmentHandler) ShowDepartment(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if department, err := d.departmentService.GetDepartmentById(ctx, id); err != nil {
		d.logger.WithContext(ctx).Error("departmentService.GetDepartmentById error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	} else {
		api.Success(ctx, department)
	}
}

// UpdateDepartment godoc
// @Summary 修改部门
// @Schemes
// @Description 修改单个部门信息
// @Tags 部门模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "部门ID"
// @Param request body api.DepartmentUpdateRequest true "params"
// @Success 200 {object} api.Response
// @Router /department/{id} [put]
func (d *DepartmentHandler) UpdateDepartment(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	req := new(api.DepartmentUpdateRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		d.logger.WithContext(ctx).Error("DepartmentHandler.UpdateDepartment error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if err := d.departmentService.UpdateDepartment(ctx, id, req); err != nil {
		d.logger.WithContext(ctx).Error("departmentService.UpdateDepartment error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, nil)
}

// DeleteDepartment godoc
// @Summary 删除部门
// @Schemes
// @Description 删除单个部门
// @Tags 部门模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "部门ID"
// @Success 200 {object} api.Response
// @Router /department/{id} [delete]
func (d *DepartmentHandler) DeleteDepartment(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := d.departmentService.DeleteDepartment(ctx, id); err != nil {
		d.logger.WithContext(ctx).Error("departmentService.DeleteDepartment error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, nil)
}
