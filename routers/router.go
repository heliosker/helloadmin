package routers

import (
	"github.com/gin-gonic/gin"
	"helloadmin/middleware"
	"helloadmin/routers/api/v1"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	apiv1.POST("/auth/login", v1.AuthLogin)
	apiv1.DELETE("/auth/logout", v1.AuthLogout)

	apiv1.Use(middleware.JWTAuthMiddleware())
	{
		apiv1.GET("/version", v1.VersionIndex)
		apiv1.POST("/version", v1.VersionStore)

		// Role
		apiv1.GET("/roles", v1.RoleIndex)
		apiv1.GET("/roles/:id", v1.RoleShow)
		apiv1.POST("/roles", v1.RoleStore)
		apiv1.PUT("/roles/:id", v1.RoleUpdate)
		apiv1.DELETE("/roles/:id", v1.RoleDestroy)

		// Administrators
		apiv1.GET("/administrators", v1.AdminIndex)
		apiv1.GET("/administrators/:id", v1.AdminShow)
		apiv1.POST("administrators", v1.AdminStore)
		apiv1.PUT("administrators/:id", v1.AdminUpdate)
		apiv1.DELETE("administrators/:id", v1.AdminDelete)

		//

		apiv1.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "test",
				"data":    "tEST",
			})
		})
	}

	return r
}
