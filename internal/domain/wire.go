//go:build wireinject
// +build wireinject

package domain

import (
	"github.com/google/wire"
	"github.com/pinguo-icc/mermaid-test/internal/infrastructure/dao"
)

func ProviderSet() wire.ProviderSet {
	return wire.NewSet(
		NewExampleRepository,
		dao.NewExampleDAO,
		wire.Bind(new(ExampleDAO), new(*dao.ExampleDAO)),

		NewModelRepository,
		dao.NewModelDAO,
		wire.Bind(new(ModelDAO), new(*dao.ModelDAO)),

		NewDeploySvc,
	)
}
