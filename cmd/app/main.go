package main

import (
	"flag"
	"fmt"
	"os"

	kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/pinguo-icc/kratos-library/v2/pdebug"
	"github.com/pinguo-icc/kratos-template-svc/internal/infrastructure/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	zapCallerSkip = 3
)

var (
	env        string // 运行环境：dev, qa, prod
	enableHTTP bool
)

func init() {
	flag.StringVar(&env, "env", "prod", "specify runtime environment: dev, qa, prod")
	flag.BoolVar(&enableHTTP, "enable-http", false, "wether enable http server")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
	var server kratos.Option
	if enableHTTP {
		server = kratos.Server(hs, gs)
	} else {
		server = kratos.Server(gs)
	}

	return kratos.New(
		kratos.Name("template"),
		kratos.Logger(logger),
		server,
	)
}

func main() {
	flag.Parse()

	cfg, err := conf.Load(env)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if env == "dev" {
		pdebug.Enable(true)
	}

	klog := newLogger(env)
	defer func() {
		_ = klog.Sync()
	}()

	logger := log.With(klog, "trace_id", tracing.TraceID())

	app, closeFunc, err := initApp(
		logger,
		cfg,
	)
	if err != nil {
		panic(err)
	}
	defer closeFunc()

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newLogger(env string) *kzap.Logger {
	var zlogger *zap.Logger
	var err error
	switch env {
	case "prod", "qa":
		cfg := zap.NewProductionConfig()
		cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		cfg.EncoderConfig.MessageKey = ""
		zlogger, err = cfg.Build(
			zap.WithCaller(false),
			zap.AddStacktrace(zapcore.FatalLevel),
			zap.AddCallerSkip(zapCallerSkip),
		)
	default:
		zlogger, err = zap.NewDevelopment(
			zap.WithCaller(false),
			zap.AddStacktrace(zapcore.FatalLevel),
			zap.AddCallerSkip(zapCallerSkip),
		)
	}

	if err != nil {
		panic(err)
	}

	return kzap.NewLogger(zlogger)
}
