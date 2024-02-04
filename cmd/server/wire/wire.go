//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"helloadmin/internal/department"
	"helloadmin/internal/login_record"
	"helloadmin/internal/menu"
	"helloadmin/internal/repository"
	"helloadmin/internal/role"
	"helloadmin/internal/server"
	"helloadmin/internal/user"
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
	role.NewRoleRepository,
	menu.NewMenuRepository,
	department.NewDeptRepository,
	login_record.NewLoginRecordRepository,
	user.NewUserRepository,
)

var serviceSet = wire.NewSet(
	role.NewRoleService,
	menu.NewMenuService,
	department.NewDepartmentService,
	login_record.NewService,
	user.NewUserService,
)

var handlerSet = wire.NewSet(
	role.NewHandler,
	menu.NewHandler,
	department.NewHandler,
	login_record.NewHandler,
	user.NewHandler,
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
		app.WithName("hello-admin-server"),
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
