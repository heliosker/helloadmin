package login_record

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helloadmin/internal/api"
	"helloadmin/pkg/log"
	"net/http"
)

type Handler struct {
	logger             *log.Logger
	loginRecordService LoginRecordService
}

func NewHandler(logger *log.Logger, svc LoginRecordService) *Handler {
	return &Handler{
		logger:             logger,
		loginRecordService: svc,
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
// @Param request query LoginRecordFindRequest true "params"
// @Success 200 {object} api.Response
// @Router /record/login [get]
func (lrh *Handler) SearchLoginRecord(ctx *gin.Context) {
	req := new(LoginRecordFindRequest)
	if err := ctx.ShouldBind(req); err != nil {
		lrh.logger.WithContext(ctx).Error("SearchLoginRecord ShouldBind error", zap.Error(err))
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if res, err := lrh.loginRecordService.Search(ctx, req); err != nil {
		api.Error(ctx, http.StatusInternalServerError, err)
	} else {
		api.Success(ctx, res)
	}
}
