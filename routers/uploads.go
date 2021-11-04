package routers

import (
	"github.com/gin-gonic/gin"
	"helloadmin/pkg/app"
	"helloadmin/pkg/errcode"
	"helloadmin/pkg/upload"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	dst := upload.GetSavePath() + upload.GetFileName(file.Filename)
	response := app.NewResponse(c)

	// 上传文件至指定目录
	if e := c.SaveUploadedFile(file, dst); e != nil {
		errRsp := errcode.UploadFileFail.WithDetails(e.Error())
		response.Error(errRsp)
		return
	}

	response.Success(gin.H{"url": dst}, app.NoMeta)
}
