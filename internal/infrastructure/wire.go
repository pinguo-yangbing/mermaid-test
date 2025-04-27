package infrastructure

import (
	"github.com/google/wire"
	pmongo "github.com/pinguo-icc/kratos-library/v2/mongo"
	"github.com/pinguo-icc/kratos-library/v2/trace"
	"github.com/pinguo-icc/kratos-template-svc/internal/infrastructure/conf"
	"github.com/pinguo-icc/kratos-template-svc/internal/infrastructure/server"
)

var ProviderSet = wire.NewSet(
	conf.ProviderSet,
	pmongo.NewDatabase,
	trace.NewTracerProvider,
	server.NewGRPCServer,
	server.NewHTTPServer,
)
