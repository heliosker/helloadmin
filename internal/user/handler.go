package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"go.uber.org/zap"
	"helloadmin/internal/api"
	"helloadmin/internal/ecode"
	logging "helloadmin/internal/login_record"
	"helloadmin/pkg/jwt"
	"helloadmin/pkg/log"
)

type Handler struct {
	log *log.Logger
	us  Service
	rs  logging.Service
}

func NewHandler(log *log.Logger, us Service, rs logging.Service) *Handler {
	return &Handler{
		log: log,
		us:  us,
		rs:  rs,
	}
}

// Search godoc
// @Summary 搜索员工
// @Schemes
// @Description 搜索员工
// @Tags 员工模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request query FindRequest true "params"
// @Success 200 {object} api.Response
// @Router /user [get]
func (h *Handler) Search(ctx *gin.Context) {
	req := new(FindRequest)
	if err := ctx.ShouldBindQuery(req); err != nil {
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if resp, err := h.us.Search(ctx, req); err != nil {
		h.log.WithContext(ctx).Error("us.Search error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
	} else {
		api.Success(ctx, resp)
	}
}

// Store godoc
// @Summary 添加员工
// @Schemes
// @Description 添加员工
// @Tags 员工模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body RegisterRequest true "params"
// @Success 200 {object} api.Response
// @Router /user [post]
func (h *Handler) Store(ctx *gin.Context) {
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
// @Summary 员工登录
// @Schemes
// @Description
// @Tags 员工模块
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
	record := logging.CreateRequest{Ip: ctx.ClientIP(), UserName: "-", Email: req.Email, Browser: browser, Platform: ua.Platform(), Os: ua.OS()}
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
// @Summary 获取员工信息
// @Schemes
// @Description
// @Tags 员工模块
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} ProfileData
// @Router /user/profile [get]
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

// UpdateProfile godoc
// @Summary 更新员工信息
// @Schemes
// @Description
// @Tags 员工模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body UpdateProfileRequest true "params"
// @Success 200 {object} api.Response
// @Router /user [put]
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
