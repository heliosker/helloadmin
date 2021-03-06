package routers

import (
	"context"
	"helloadmin/app/service"
	"helloadmin/pkg/app"
	"helloadmin/pkg/errcode"
	"helloadmin/pkg/upload"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	dst := upload.SavePath() + upload.GetFileName(file.Filename)
	response := app.NewResponse(c)
	// 上传文件至指定目录
	if e := c.SaveUploadedFile(file, dst); e != nil {
		errRsp := errcode.UploadFileFail.WithDetails(e.Error())
		response.Error(errRsp)
		return
	}

	response.Success(gin.H{"url": dst}, app.NoMeta)
}

func (u Upload) UploadQiniuOss(c *gin.Context) {

	svc := service.New(c)
	var (
		accessKey = svc.GetConfigValue("QINIU_ACCESS_KEY")
		secretKey = svc.GetConfigValue("QINIU_SECRET_KEY")
		bucket    = svc.GetConfigValue("QINIU_BUCKET")
	)
	file, _ := c.FormFile("file")
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		//Params: map[string]string{
		//	"x:name": "github logo",
		//},
	}
	data, _ := file.Open()
	key := upload.SavePath() + strconv.FormatInt(time.Now().Unix(), 10)
	err := formUploader.Put(context.Background(), &ret, upToken, key, data, file.Size, &putExtra)
	rsp := app.NewResponse(c)
	if err != nil {
		rsp.Error(errcode.UploadFileFail.WithDetails(err.Error()))
		return
	}
	rsp.Success(gin.H{"key": ret.Key, "hash": ret.Hash, "host": "http://oss.helloadmin.cn"}, app.NoMeta)
}
