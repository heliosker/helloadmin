//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"helloadmin/internal/handler"
	"helloadmin/internal/repository"
	"helloadmin/internal/server"
	"helloadmin/internal/service"
	"helloadmin/pkg/app"
	"helloadmin/pkg/helper/sid"
	"helloadmin/pkg/jwt"
	"helloadmin/pkg/log"
	"helloadmin/pkg/server/http"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewRoleRepository,
	repository.NewMenuRepository,
	repository.NewDepartmentRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewRoleService,
	service.NewMenuService,
	service.NewDepartmentService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewRoleHandler,
	handler.NewMenuHandler,
	handler.NewDepartmentHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
	server.NewTask,
)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {

	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
