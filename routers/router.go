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
	apiv1.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"Test": "test action 4"})
	})
	apiv1.POST("/auth/login", v1.AuthLogin)
	apiv1.DELETE("/auth/logout", v1.AuthLogout)

	role := v1.NewRole()
	ver := v1.NewVersion()
	menu := v1.NewMenu()
	apiv1.Use(middleware.JWT())
	{
		// Upload file
		apiv1.POST("/upload/qiniu", NewUpload().UploadQiniuOss)
		apiv1.POST("/upload", NewUpload().UploadFile)
		apiv1.GET("/me", v1.AdminMe)

		apiv1.GET("/version", ver.Index)
		apiv1.POST("/version", ver.Store)

		// Role
		apiv1.GET("/roles", role.Index)
		apiv1.GET("/roles/:id", role.Show)
		apiv1.POST("/roles", role.Store)
		apiv1.PUT("/roles/:id", role.Update)
		apiv1.POST("/roles/:id/menus", role.Update)
		apiv1.PUT("/roles/:id/menus", role.Update)
		apiv1.DELETE("/roles/:id", role.Destroy)

		// Administrators
		apiv1.GET("/administrators", v1.AdminIndex)
		apiv1.GET("/administrators/:id", v1.AdminShow)
		apiv1.POST("/administrators", v1.AdminStore)
		apiv1.PUT("/administrators/:id", v1.AdminUpdate)
		apiv1.DELETE("/administrators/:id", v1.AdminDelete)

		// Menu
		apiv1.GET("/menus", menu.Index)
		apiv1.GET("/menus/:id", menu.Show)
		apiv1.POST("/menus", menu.Store)
		apiv1.PUT("/menus/:id", menu.Update)
		apiv1.DELETE("/menus/:id", menu.Delete)

	}

	return r
}
