package domain

import (
	"github.com/pinguo-icc/kratos-library/v2/base"
	"github.com/pinguo-icc/mermaid-test/internal/infrastructure/dao"
)

type Example struct {
	*dao.Example
}

// exampleDO2DAO 将领域对象转换为dao对象，根据需要可以重新实现
func exampleDO2DAO(u *Example) (base.IBaseDO, error) {
	if u == nil {
		return nil, nil
	}

	return u.Example, nil
}

// exampleDAO2DO 将DAO对象转换为领域对象，根据需要可以重新实现
func exampleDAO2DO(u *dao.Example) (*Example, error) {
	if u == nil {
		return nil, nil
	}

	return &Example{
		Example: u,
	}, nil
}
