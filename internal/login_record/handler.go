package login_record

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helloadmin/internal/api"
	"helloadmin/pkg/log"
	"net/http"
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

// SearchLoginRecord godoc
// @Summary 登录日志
// @Schemes
// @Description 登录日志列表
// @Tags 日志模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request query FindRequest true "params"
// @Success 200 {object} api.Response
// @Router /record/login [get]
func (lrh *Handler) SearchLoginRecord(ctx *gin.Context) {
	req := new(FindRequest)
	if err := ctx.ShouldBind(req); err != nil {
		lrh.log.WithContext(ctx).Error("SearchLoginRecord ShouldBind error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if res, err := lrh.svc.Search(ctx, req); err != nil {
		api.Error(ctx, http.StatusInternalServerError, err)
	} else {
		api.Success(ctx, res)
	}
}
