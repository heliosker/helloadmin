package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"go.uber.org/zap"
	"helloadmin/internal/api"
	"helloadmin/internal/ecode"
	login_log "helloadmin/internal/login_record"
	"helloadmin/pkg/jwt"
	"helloadmin/pkg/log"
	"net/http"
)

type Handler struct {
	log *log.Logger
	us  Service
	rs  login_log.Service
}

func NewHandler(logger *log.Logger, us Service, rs login_log.Service) *Handler {
	return &Handler{
		log: logger,
		us:  us,
		rs:  rs,
	}
}

// Register godoc
// @Summary 用户注册
// @Schemes
// @Description 目前只支持邮箱登录
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "params"
// @Success 200 {object} api.Response
// @Router /register [post]
func (h *Handler) Register(ctx *gin.Context) {
	req := new(RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if err := h.us.Register(ctx, req); err != nil {
		h.log.WithContext(ctx).Error("us.Register error", zap.Error(err))
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
// @Param request body LoginRequest true "params"
// @Success 200 {object} LoginResponse
// @Router /login [post]
func (h *Handler) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	ua := user_agent.New(ctx.Request.UserAgent())
	browser, _ := ua.Browser()
	record := login_log.LoginRecordRequest{Ip: ctx.ClientIP(), UserName: "-", Email: req.Email, Browser: browser, Platform: ua.Platform(), Os: ua.OS()}
	resp, err := h.us.Login(ctx, &req)
	if err != nil {
		record.ErrorMessage = err.Error()
		_ = h.rs.Create(ctx, &record)
		api.Error(ctx, http.StatusUnauthorized, err)
		return
	}
	_ = h.rs.Create(ctx, &record)
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
// @Success 200 {object} GetProfileResponseData
// @Router /user [get]
func (h *Handler) GetProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		api.Error(ctx, http.StatusUnauthorized, ecode.ErrUnauthorized)
		return
	}
	user, err := h.us.GetProfile(ctx, userId)
	if err != nil {
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	api.Success(ctx, user)
}

func (h *Handler) UpdateProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	var req UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if err := h.us.UpdateProfile(ctx, userId, &req); err != nil {
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, nil)
}

func GetUserIdFromCtx(ctx *gin.Context) string {
	v, exists := ctx.Get("claims")
	if !exists {
		return ""
	}
	return v.(*jwt.MyCustomClaims).UserId
}
