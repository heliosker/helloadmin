package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"helloadmin/docs"
)

func SwaggerInit() gin.HandlerFunc {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "HelloAdmin API"
	docs.SwaggerInfo.Description = "HelloAdmin 接口文档"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "http://localhost:9010"
	docs.SwaggerInfo.BasePath = "/api/v1/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}
