package routers

import (
	"github.com/gin-gonic/gin"
	"helloadmin/middleware"
	"helloadmin/routers/api"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	apiv1.POST("/auth", api.AuthLogin)

	apiv1.Use(middleware.JWTAuthMiddleware())
	{

		apiv1.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "test",
				"data":"tEST",
			})
		})
	}

	return r
}
