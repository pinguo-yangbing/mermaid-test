package dao

import (
	"time"

	"github.com/pinguo-icc/kratos-library/v2/base"
	pmongo "github.com/pinguo-icc/kratos-library/v2/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MirrorVendorType 镜像供应商类型
type MirrorVendorType string

const (
	MirrorVendorTypeAWS MirrorVendorType = "aws"
	MirrorVendorTypeAli MirrorVendorType = "ali"
)

// CodeRepositoryType 代码仓库类型
type CodeRepositoryType string

const (
	CodeRepositoryTypeGit CodeRepositoryType = "git"
)

// BuildingStatus 构建状态
type BuildingStatus string

const (
	BuildingStatusBuilding  BuildingStatus = "building"
	BuildingStatusSucceeded BuildingStatus = "succeeded"
	BuildingStatusFailed    BuildingStatus = "failed"
)

// CallType 调用类型
type CallType string

const (
	CallTypeSync  CallType = "sync"
	CallTypeAsync CallType = "async"
)

// ReplicaConfigType 副本配置类型
type ReplicaConfigType string

const (
	ReplicaConfigTypeAuto   ReplicaConfigType = "auto"
	ReplicaConfigTypeManual ReplicaConfigType = "manual"
)

// DeployConfigType 部署配置类型
type DeployConfigType string

const (
	DeployConfigTypeBase   DeployConfigType = "base"
	DeployConfigTypeCustom DeployConfigType = "custom"
)

// Status 状态
type Status string

const (
	StatusNormal  Status = "normal"
	StatusDeleted Status = "deleted"
)

// MirrorVendor 镜像供应商
type MirrorVendor struct {
	Name      string           `bson:"name"`
	Region    string           `bson:"region"`
	AcrRepoID string           `bson:"acrRepoId"`
	Type      MirrorVendorType `bson:"type"`
}

// CodeRepository 代码仓库
type CodeRepository struct {
	Name      string             `bson:"name"`
	Region    string             `bson:"region"`
	AcrRepoID string             `bson:"acrRepoId"`
	Type      CodeRepositoryType `bson:"type"`
}

// Auto 自动副本配置
type Auto struct {
	MinReplicas int   `bson:"minReplicas"`
	MaxReplicas int   `bson:"maxReplicas"`
	MaxDelay    int64 `bson:"maxDelay"`
}

// Manual 手动副本配置
type Manual struct {
	Replicas int `bson:"replicas"`
}

// ReplicaConfig 副本配置
type ReplicaConfig struct {
	Type   ReplicaConfigType `bson:"type"`
	Auto   *Auto             `bson:"auto,omitempty"`
	Manual *Manual           `bson:"manual,omitempty"`
}

// DeployConfigBase 基础部署配置
type DeployConfigBase struct {
	Envs             map[string]string `bson:"envs"`
	CPURequest       string            `bson:"cpuRequest"`
	CPULimit         string            `bson:"cpuLimit"`
	MemoryRequest    string            `bson:"memoryRequest"`
	MemoryLimit      string            `bson:"memoryLimit"`
	GPURequest       string            `bson:"gpuRequest"`
	GPUMemoryRequest string            `bson:"gpuMemoryRequest"`
	MaxStartTime     int               `bson:"maxStartTime"`
	Clusters         []string          `bson:"clusters"`
}

// DeployConfig 部署配置
type DeployConfig struct {
	Type              DeployConfigType  `bson:"type"`
	Base              *DeployConfigBase `bson:"base,omitempty"`
	CustomDeployYaml  string            `bson:"customDeployYaml,omitempty"`
	CustomKarmadaYaml string            `bson:"customKarmadaYaml,omitempty"`
}

// TrafficConfig 流量配置
type TrafficConfig struct {
	Weight int  `bson:"weight"`
	Dev    bool `bson:"dev"`
}

// Image 镜像
type Image struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	ModelID        string             `bson:"modelId"`
	Tag            string             `bson:"tag"`
	Comment        string             `bson:"comment"`
	FlowPercentage uint32             `bson:"flowPercentage"`
	ImageURI       string             `bson:"imageURI"`
	Status         BuildingStatus     `bson:"status"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
}

// Deployment 部署
type Deployment struct {
	Template      string         `bson:"template"`
	DeployID      string         `bson:"deployId"`
	ModelID       string         `bson:"modelId"`
	Tag           string         `bson:"tag"`
	CallType      CallType       `bson:"callType"`
	ReplicaConfig *ReplicaConfig `bson:"replicaConfig"`
	Config        *DeployConfig  `bson:"config"`
	Status        Status         `bson:"status"`
	CreatedAt     time.Time      `bson:"createdAt"`
	UpdatedAt     time.Time      `bson:"updatedAt"`
	ApplyTime     time.Time      `bson:"applyTime"`
	TrafficConfig *TrafficConfig `bson:"trafficConfig"`
}

// Model 模型实体
type Model struct {
	base.BaseDO    `bson:"inline"`
	Name           string          `bson:"name"`
	Description    string          `bson:"description"`
	Creator        string          `bson:"creator"`
	Mender         string          `bson:"mender"`
	CreatedAt      time.Time       `bson:"createdAt"`
	UpdatedAt      time.Time       `bson:"updatedAt"`
	MirrorVendor   *MirrorVendor   `bson:"mirrorVendor"`
	CodeRepository *CodeRepository `bson:"codeRepository"`
	Images         []*Image        `bson:"images"`
	Deployments    []*Deployment   `bson:"deployments"`
}

// ModelDAO Model数据访问对象
type ModelDAO struct {
	*base.BaseDAOImpl[Model]
}

// NewModelDAO 创建ModelDAO对象
func NewModelDAO(db *mongo.Database) (*ModelDAO, error) {
	dao := &ModelDAO{
		BaseDAOImpl: &base.BaseDAOImpl[Model]{
			Coll: pmongo.NewCollectionSelector(db, "model"),
		},
	}

	return dao, nil
}
