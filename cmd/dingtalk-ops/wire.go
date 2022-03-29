//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/rogerogers/dingtalk-ops/internal/biz"
	"github.com/rogerogers/dingtalk-ops/internal/conf"
	"github.com/rogerogers/dingtalk-ops/internal/data"
	"github.com/rogerogers/dingtalk-ops/internal/server"
	"github.com/rogerogers/dingtalk-ops/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Dingtalk, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
