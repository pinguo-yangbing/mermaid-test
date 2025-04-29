package application

import (
	"github.com/pinguo-icc/mermaid-test/api"
)

type App struct {
	api.UnsafeTemplateServer
}

var _ api.TemplateServer = (*App)(nil)
