package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"helloadmin/internal/ecode"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Pagination struct {
	Page  int `json:"page" binding:"required,min=1" example:"1"`
	Size  int `json:"size" binding:"required,min=1,max=100" example:"10"`
	Count int `json:"count"`
}

func Success(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	e := ecode.ErrSuccess
	resp := Response{Code: e.Int(), Message: e.Error(), Data: data}
	ctx.JSON(http.StatusOK, resp)
}

func Error(ctx *gin.Context, httpCode int, err error) {
	var c ecode.ErrCode
	var m string
	switch {
	case errors.As(err, &c):
		c = c.Code()
		m = c.String()
	default:
		c = ecode.ErrInternalServerError
		m = err.Error()
	}
	resp := Response{Code: c.Int(), Message: m}
	ctx.JSON(httpCode, resp)
}
