package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"helloadmin/api"
	"helloadmin/docs"
	"helloadmin/internal/handler"
	"helloadmin/internal/middleware"
	"helloadmin/pkg/jwt"
	"helloadmin/pkg/log"
	"helloadmin/pkg/server/http"
)

func NewHTTPServer(
	logger *log.Logger,
	cfg *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
	roleHandler *handler.RoleHandler,
	menuHandler *handler.MenuHandler,
	departHandler *handler.DepartmentHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(cfg.GetString("http.host")),
		http.WithServerPort(cfg.GetInt("http.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/api"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", cfg.GetInt("http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)
	s.GET("/", func(ctx *gin.Context) {
		api.Success(ctx, map[string]interface{}{
			":)": "Thank you for using HelloAdmin!",
		})
	})

	group := s.Group("/api")
	{
		// No route group has permission
		noAuthRouter := group.Group("/")
		{
			noAuthRouter.POST("/register", userHandler.Register)
			noAuthRouter.POST("/login", userHandler.Login)
		}
		// Non-strict permission routing group
		noStrictAuthRouter := group.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
		{
			noStrictAuthRouter.GET("/user", userHandler.GetProfile)
		}

		// Strict permission routing group
		strictAuthRouter := group.Group("/").Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
		}

		role := group.Group("/role").Use(middleware.StrictAuth(jwt, logger))
		{
			role.GET("", roleHandler.GetRole)
			role.POST("", roleHandler.StoreRole)
			role.GET("/:id", roleHandler.ShowRole)
			role.PUT("/:id", roleHandler.UpdateRole)
			role.DELETE("/:id", roleHandler.DeleteRole)
		}

		menu := group.Group("/menu").Use(middleware.StrictAuth(jwt, logger))
		{
			menu.GET("", menuHandler.GetMenu)
			menu.POST("", menuHandler.StoreMenu)
			menu.GET("/:id", menuHandler.ShowMenu)
			menu.PUT("/:id", menuHandler.UpdateMenu)
			menu.DELETE("/:id", menuHandler.DeleteMenu)
		}
		depart := group.Group("department").Use(middleware.StrictAuth(jwt, logger))
		{
			depart.GET("", departHandler.GetDepartment)
			depart.POST("", departHandler.StoreDepartment)
			depart.GET("/:id", departHandler.ShowDepartment)
			depart.PUT("/:id", departHandler.UpdateDepartment)
			depart.DELETE("/:id", departHandler.DeleteDepartment)
		}

	}

	return s
}
