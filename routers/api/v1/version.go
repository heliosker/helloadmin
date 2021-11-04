package v1

import (
	"github.com/gin-gonic/gin"
	"helloadmin/models"
	"helloadmin/pkg/app"
	"helloadmin/pkg/errcode"
	"helloadmin/pkg/utils"
	"strconv"
)

type Version struct {
}

func NewVersion() Version {
	return Version{}
}

func (v Version) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	var count int64
	var version []models.Version
	models.DB.Model(&version).Count(&count)
	ret := models.DB.Scopes(utils.Paginate(page, size)).Find(&version)
	if ret.Error != nil {
		app.NewResponse(c).Error(errcode.CreatedFail)
		return
	}
	app.NewResponse(c).Success(version, count)
}

func (v Version) Store(c *gin.Context) {
	var version models.Version
	_ = c.ShouldBindJSON(&version)

	err := models.DB.Create(&version).Error
	if err != nil {
		app.NewResponse(c).Error(errcode.CreatedFail.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).Success(version, app.NoMeta)
}
