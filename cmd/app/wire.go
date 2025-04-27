//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/pinguo-icc/kratos-template-svc/internal/application"
	_ "github.com/pinguo-icc/kratos-template-svc/internal/domain"
	infra "github.com/pinguo-icc/kratos-template-svc/internal/infrastructure"
	"github.com/pinguo-icc/kratos-template-svc/internal/infrastructure/conf"
)

//initApp init kratos application.
//go:generate kratos t wire
func initApp(
	log.Logger,
	*conf.Bootstrap,
) (*kratos.App, func(), error) {
	panic(wire.Build(
		infra.ProviderSet,
		application.ProviderSet,
		// domain.ProviderSet,
		newApp,
	))
}
