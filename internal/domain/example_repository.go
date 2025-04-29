package domain

import (
	"github.com/pinguo-icc/kratos-library/v2/base"
	"github.com/pinguo-icc/mermaid-test/internal/infrastructure/dao"
)

var (
	_ ExampleDAO        = (*dao.ExampleDAO)(nil)
	_ ExampleRepository = (*exampleRepository)(nil)
)

type (
	ExampleDAO interface {
		base.BaseDAO[dao.Example]
	}

	ExampleRepository interface {
		base.BaseRepository[Example]
	}
)

type exampleRepository struct {
	utDAO ExampleDAO
	base.BaseRepositoryImpl[Example, dao.Example]
}

// TODO: 在此处为exampleRepository添加其它成员方法，或者覆写base.base.BaseRepositoryImpl已实现的方法

// NewExampleRepository 创建ExampleRepository 仓储层对象
func NewExampleRepository(d ExampleDAO) (ExampleRepository, error) {
	utr := &exampleRepository{
		utDAO: d,
	}

	utr.BaseRepositoryImpl = base.BaseRepositoryImpl[Example, dao.Example]{
		BDAO:    d,
		FuncT2D: exampleDO2DAO,
		FuncD2T: exampleDAO2DO,
	}

	return utr, nil
}
