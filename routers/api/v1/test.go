package v1

import (
	"github.com/gin-gonic/gin"
	"helloadmin/pkg/app"
	"time"
)

func Test(c *gin.Context) {
	app.NewResponse(c).Success(gin.H{"now": time.Now().Format("2006-01-02 15:04:05")}, app.NoMeta)
}
