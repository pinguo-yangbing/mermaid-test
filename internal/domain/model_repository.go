package domain

import (
	"github.com/pinguo-icc/kratos-library/v2/base"
	"github.com/pinguo-icc/mermaid-test/internal/infrastructure/dao"
)

var (
	_ ModelDAO        = (*dao.ModelDAO)(nil)
	_ ModelRepository = (*modelRepository)(nil)
)

type (
	ModelDAO interface {
		base.BaseDAO[dao.Model]
	}

	ModelRepository interface {
		base.BaseRepository[Model]
		AppendVersion(id string, image *dao.Image) error
		AppendDeploy(id string, deployment *dao.Deployment) error
	}
)

type modelRepository struct {
	modelDAO ModelDAO
	base.BaseRepositoryImpl[Model, dao.Model]
}

// AppendVersion 添加版本
func (r *modelRepository) AppendVersion(id string, image *dao.Image) error {
	// 实现添加版本逻辑
	return nil
}

// AppendDeploy 添加部署
func (r *modelRepository) AppendDeploy(id string, deployment *dao.Deployment) error {
	// 实现添加部署逻辑
	return nil
}

// NewModelRepository 创建ModelRepository仓储层对象
func NewModelRepository(d ModelDAO) (ModelRepository, error) {
	repo := &modelRepository{
		modelDAO: d,
	}

	repo.BaseRepositoryImpl = base.BaseRepositoryImpl[Model, dao.Model]{
		BDAO:    d,
		FuncT2D: modelDO2DAO,
		FuncD2T: modelDAO2DO,
	}

	return repo, nil
}
