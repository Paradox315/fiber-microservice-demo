//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"fiber-demo/app/demo/internal/biz"
	"fiber-demo/app/demo/internal/conf"
	"fiber-demo/app/demo/internal/data"
	"fiber-demo/app/demo/internal/server"
	"fiber-demo/app/demo/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Logger, *conf.Storage) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
