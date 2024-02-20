package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"helloadmin/docs"
	"helloadmin/internal/api"
	"helloadmin/internal/department"
	"helloadmin/internal/login_record"
	"helloadmin/internal/menu"
	"helloadmin/internal/middleware"
	"helloadmin/internal/role"
	"helloadmin/internal/user"
	"helloadmin/pkg/jwt"
	"helloadmin/pkg/log"
	"helloadmin/pkg/server/http"
)

func NewHTTPServer(
	logger *log.Logger,
	cfg *viper.Viper,
	jwt *jwt.JWT,
	userHandler *user.Handler,
	roleHandler *role.Handler,
	menuHandler *menu.Handler,
	departHandler *department.Handler,
	loginRecordHandler *login_record.Handler,
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
		// ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", cfg.GetInt("http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		// middleware.SignMiddleware(log),
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
			noAuthRouter.POST("/login", userHandler.Login)
		}
		// Strict permission routing group
		userRouter := group.Group("/user").Use(middleware.StrictAuth(jwt, logger))
		{
			userRouter.GET("", userHandler.Search)
			userRouter.GET("/:id", userHandler.Show)
			userRouter.GET("/profile", userHandler.GetProfile)
			userRouter.POST("", userHandler.Store)
			userRouter.PUT("/:id", userHandler.Update)
		}

		roleRouter := group.Group("/role").Use(middleware.StrictAuth(jwt, logger))
		{
			roleRouter.GET("", roleHandler.GetRole)
			roleRouter.POST("", roleHandler.StoreRole)
			roleRouter.GET("/:id", roleHandler.ShowRole)
			roleRouter.PUT("/:id", roleHandler.UpdateRole)
			roleRouter.PUT("/:id/menu", roleHandler.UpdateRoleMenu)
			roleRouter.DELETE("/:id", roleHandler.DeleteRole)
		}

		menuRouter := group.Group("/menu").Use(middleware.StrictAuth(jwt, logger))
		{
			menuRouter.GET("", menuHandler.GetMenu)
			menuRouter.GET("/option", menuHandler.GetOption)
			menuRouter.POST("", menuHandler.StoreMenu)
			menuRouter.GET("/:id", menuHandler.ShowMenu)
			menuRouter.PUT("/:id", menuHandler.UpdateMenu)
			menuRouter.DELETE("/:id", menuHandler.DeleteMenu)
		}
		deptRouter := group.Group("department").Use(middleware.StrictAuth(jwt, logger))
		{
			deptRouter.GET("", departHandler.GetDepartment)
			deptRouter.POST("", departHandler.StoreDepartment)
			deptRouter.GET("/:id", departHandler.ShowDepartment)
			deptRouter.PUT("/:id", departHandler.UpdateDepartment)
			deptRouter.DELETE("/:id", departHandler.DeleteDepartment)
		}
		recordRouter := group.Group("/record").Use(middleware.StrictAuth(jwt, logger))
		{
			recordRouter.GET("login", loginRecordHandler.SearchLoginRecord)
			recordRouter.GET("operation", loginRecordHandler.SearchLoginRecord)
		}

	}

	return s
}
