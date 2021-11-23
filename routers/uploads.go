package routers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"helloadmin/config"
	"helloadmin/pkg/app"
	"helloadmin/pkg/errcode"
	"helloadmin/pkg/upload"
	"strconv"
	"time"
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

	config, err := config.Load().GetSection("app")
	var (
		accessKey = config.Key("QINIU_ACCESS_KEY").String()
		secretKey = config.Key("QINIU_SECRET_KEY").String()
		bucket    = config.Key("QINIU_TEST_BUCKET").String()
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
	err = formUploader.Put(context.Background(), &ret, upToken, key, data, file.Size, &putExtra)
	rsp := app.NewResponse(c)
	if err != nil {
		fmt.Println(err)
		rsp.Error(errcode.UploadFileFail.WithDetails(err.Error()))
		return
	}
	//fmt.Println(ret.Key, ret.Hash)
	rsp.Success(gin.H{"key": ret.Key, "hash": ret.Hash, "host": "http://oss.helloadmin.cn"}, app.NoMeta)

}
