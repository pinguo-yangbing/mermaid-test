package domain

import (
	"github.com/pinguo-icc/kratos-library/v2/base"
	"github.com/pinguo-icc/mermaid-test/internal/infrastructure/dao"
)

// Model 模型领域实体
type Model struct {
	*dao.Model
}

// Validate 验证模型
func (m *Model) Validate() error {
	// 实现验证逻辑
	return nil
}

// modelDO2DAO 将领域对象转换为dao对象
func modelDO2DAO(m *Model) (base.IBaseDO, error) {
	if m == nil {
		return nil, nil
	}

	return m.Model, nil
}

// modelDAO2DO 将DAO对象转换为领域对象
func modelDAO2DO(m *dao.Model) (*Model, error) {
	if m == nil {
		return nil, nil
	}

	return &Model{
		Model: m,
	}, nil
}
