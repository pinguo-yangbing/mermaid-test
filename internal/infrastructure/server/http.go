package server

import (
	klog "github.com/pinguo-icc/kratos-library/v2/log"
	"github.com/pinguo-icc/kratos-template-svc/api"
	"github.com/pinguo-icc/kratos-template-svc/internal/infrastructure/conf"
	"go.opentelemetry.io/otel/trace"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a http server.
func NewHTTPServer(
	c *conf.HTTP, logger log.Logger,
	svc api.TemplateServer, provider trace.TracerProvider,
) *http.Server {
	opts := []http.ServerOption{
		http.Address(c.Address),
		http.Timeout(c.Timeout),
		http.Logger(logger),
		http.Middleware(
			recovery.Recovery(recovery.WithLogger(logger)),
			tracing.Server(tracing.WithTracerProvider(provider)), // 顺序有要求, 需要在logger前, 否则logger不能获取span/trace id
			klog.ServerMiddleware(logger),
		),
	}
	srv := http.NewServer(opts...)

	srv.Route("/_invoke").POST("/{name}", invoke(svc))

	return srv
}
