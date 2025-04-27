package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	kierr "github.com/pinguo-icc/kratos-library/v2/ierr"
	klog "github.com/pinguo-icc/kratos-library/v2/log"
	"github.com/pinguo-icc/kratos-template-svc/api"
	"github.com/pinguo-icc/kratos-template-svc/internal/infrastructure/conf"
	"go.opentelemetry.io/otel/trace"
)

// NewGRPCServer new a grpc server.
func NewGRPCServer(
	c *conf.GRPC, logger log.Logger,
	svc api.TemplateServer, provider trace.TracerProvider,
) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.Address(c.Address),
		grpc.Timeout(c.Timeout),
		grpc.Logger(logger),
		grpc.Middleware(
			recovery.Recovery(recovery.WithLogger(logger)),
			tracing.Server(tracing.WithTracerProvider(provider)), // 顺序有要求, 需要在logger前, 否则logger不能获取span/trace id
			kierr.GRPCServerMiddleware(),                         // 顺序要求，需要在记录日志后才将响应错误转换为grpc的错误
			klog.ServerMiddleware(logger),
		),
	}

	srv := grpc.NewServer(opts...)
	api.RegisterTemplateServer(srv, svc)

	return srv
}
