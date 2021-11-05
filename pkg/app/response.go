package app

import (
	"github.com/gin-gonic/gin"
	"helloadmin/pkg/errcode"
	"net/http"
)

const NoMeta = -1

type Response struct {
	Ctx *gin.Context
}

type Meta struct {
	Page  int   `json:"page"`
	Size  int   `json:"size"`
	Count int64 `json:"count"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) Success(data interface{}, totalRows int64) {
	if data == nil {
		data = gin.H{}
	}
	if totalRows > NoMeta {
		// Items
		r.Ctx.JSON(http.StatusOK, gin.H{
			"message": errcode.Success.Message(),
			"code":    errcode.Success.Code(),
			"data":    data,
			"meta": Meta{
				Page:  GetPage(r.Ctx),
				Size:  GetPageSize(r.Ctx),
				Count: totalRows,
			},
		})
		return
	}

	// No pagination
	r.Ctx.JSON(http.StatusOK, gin.H{
		"message": errcode.Success.Message(),
		"code":    errcode.Success.Code(),
		"data":    data,
	})
}

func (r *Response) Error(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "message": err.Message()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
