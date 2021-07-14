// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/martinv13/go-shiny/controllers"
	"github.com/martinv13/go-shiny/models"
	"github.com/martinv13/go-shiny/modules/appserver"
	"github.com/martinv13/go-shiny/modules/config"
	"github.com/martinv13/go-shiny/modules/ssehandler"
	"github.com/martinv13/go-shiny/modules/vfsdata"
	"github.com/martinv13/go-shiny/server"
)

func InitializeServer(flags config.RunFlags) (*server.AppRouter, error) {
	wire.Build(server.NewAppRouter, models.NewDB, vfsdata.NewStaticPaths,
		ssehandler.NewMessageBroker, appserver.NewAppServer,
		config.NewConfigViper, wire.Bind(new(config.Config), new(*config.ConfigViper)),
		models.NewAppModelDB, wire.Bind(new(models.AppModel), new(*models.AppModelDB)),
		models.NewUserModelDB, wire.Bind(new(models.UserModel), new(*models.UserModelDB)),
		models.NewGroupModelDB, wire.Bind(new(models.GroupModel), new(*models.GroupModelDB)),
		controllers.NewAppController, controllers.NewUserController, controllers.NewGroupController,
		controllers.NewAuthController)
	return &server.AppRouter{}, nil
}
