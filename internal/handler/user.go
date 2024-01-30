package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helloadmin/api"
	"helloadmin/internal/ecode"
	"helloadmin/internal/service"
	"net/http"
)

type UserHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

// Register godoc
// @Summary 用户注册
// @Schemes
// @Description 目前只支持邮箱登录
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body api.RegisterRequest true "params"
// @Success 200 {object} api.Response
// @Router /register [post]
func (h *UserHandler) Register(ctx *gin.Context) {
	req := new(api.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		api.Error(ctx, http.StatusBadRequest, ecode.ErrBadRequest)
		return
	}
	if err := h.userService.Register(ctx, req); err != nil {
		h.logger.WithContext(ctx).Error("userService.Register error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, nil)
}

// Login godoc
// @Summary 账号登录
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body api.LoginRequest true "params"
// @Success 200 {object} api.LoginResponse
// @Router /login [post]
func (h *UserHandler) Login(ctx *gin.Context) {
	var req api.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	resp, err := h.userService.Login(ctx, &req)
	if err != nil {
		api.Error(ctx, http.StatusUnauthorized, err)
		return
	}
	api.Success(ctx, resp)
}

// GetProfile godoc
// @Summary 获取用户信息
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} api.GetProfileResponseData
// @Router /user [get]
func (h *UserHandler) GetProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		api.Error(ctx, http.StatusUnauthorized, ecode.ErrUnauthorized)
		return
	}
	user, err := h.userService.GetProfile(ctx, userId)
	if err != nil {
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	api.Success(ctx, user)
}

func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	var req api.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if err := h.userService.UpdateProfile(ctx, userId, &req); err != nil {
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, nil)
}
