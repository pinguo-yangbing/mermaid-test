package application

import (
	"github.com/pinguo-icc/kratos-template-svc/api"
)

type App struct {
	api.UnsafeTemplateServer
}

var _ api.TemplateServer = (*App)(nil)
