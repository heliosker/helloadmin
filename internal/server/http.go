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
		usr := group.Group("/user").Use(middleware.StrictAuth(jwt, logger))
		{
			usr.GET("", userHandler.Search)
			usr.GET("/:id", userHandler.Show)
			usr.GET("/profile", userHandler.GetProfile)
			usr.POST("", userHandler.Store)
			usr.PUT("/:id", userHandler.Update)
			usr.DELETE("/:id", userHandler.Delete)
		}

		ror := group.Group("/role").Use(middleware.StrictAuth(jwt, logger))
		{
			ror.GET("", roleHandler.GetRole)
			ror.POST("", roleHandler.StoreRole)
			ror.GET("/:id", roleHandler.ShowRole)
			ror.PUT("/:id", roleHandler.UpdateRole)
			ror.PUT("/:id/menu", roleHandler.UpdateRoleMenu)
			ror.DELETE("/:id", roleHandler.DeleteRole)
		}

		mer := group.Group("/menu").Use(middleware.StrictAuth(jwt, logger))
		{
			mer.GET("", menuHandler.GetMenu)
			mer.GET("/option", menuHandler.GetOption)
			mer.POST("", menuHandler.StoreMenu)
			mer.GET("/:id", menuHandler.ShowMenu)
			mer.PUT("/:id", menuHandler.UpdateMenu)
			mer.DELETE("/:id", menuHandler.DeleteMenu)
		}
		der := group.Group("department").Use(middleware.StrictAuth(jwt, logger))
		{
			der.GET("", departHandler.GetDepartment)
			der.POST("", departHandler.StoreDepartment)
			der.GET("/:id", departHandler.ShowDepartment)
			der.PUT("/:id", departHandler.UpdateDepartment)
			der.DELETE("/:id", departHandler.DeleteDepartment)
		}
		rer := group.Group("/record").Use(middleware.StrictAuth(jwt, logger))
		{
			rer.GET("login", loginRecordHandler.SearchLoginRecord)
			rer.GET("operation", loginRecordHandler.SearchLoginRecord)
		}

	}

	return s
}
