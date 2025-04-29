package application

import (
	"github.com/google/wire"
	"github.com/pinguo-icc/mermaid-test/api"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(App), "*"),

	wire.Bind(new(api.TemplateServer), new(*App)),
	wire.InterfaceValue(new(api.UnsafeTemplateServer), new(api.UnimplementedTemplateServer)),
)
