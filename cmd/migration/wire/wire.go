//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"helloadmin/internal/repository"
	"helloadmin/internal/server"
	"helloadmin/internal/user"
	"helloadmin/pkg/app"
	"helloadmin/pkg/log"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	user.NewUserRepository,
)

// build App
func newApp(migrate *server.Migrate) *app.App {
	return app.NewApp(
		app.WithServer(migrate),
		app.WithName("demo-migrate"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		server.NewMigrate,
		newApp,
	))
}
