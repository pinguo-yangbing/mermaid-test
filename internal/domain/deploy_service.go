package domain

// DeploySvc 部署服务领域服务
type DeploySvc struct {
	Karmada         Karmada
	modelRepository ModelRepository
}

// Deploy 部署模型
func (s *DeploySvc) Deploy(modelID string, tag string, deployConfig interface{}) error {
	// 部署模型实现
	return nil
}

// ApplyModel 应用模型
func (s *DeploySvc) ApplyModel(modelID string, tag string) error {
	// 应用模型实现
	return nil
}

// Clusters 获取集群列表
func (s *DeploySvc) Clusters() ([]string, error) {
	return s.Karmada.GetClusters()
}

// ResourceSummary 获取资源摘要
func (s *DeploySvc) ResourceSummary() (map[string]interface{}, error) {
	return s.Karmada.GetResourceSummary()
}

// DeploymentsByModelTag 根据模型标签获取部署
func (s *DeploySvc) DeploymentsByModelTag(modelTag string) ([]interface{}, error) {
	return s.Karmada.GetDeploymentsByModelTag(modelTag)
}

// NewDeploySvc 创建部署服务
func NewDeploySvc(k Karmada, mr ModelRepository) (*DeploySvc, error) {
	return &DeploySvc{
		Karmada:         k,
		modelRepository: mr,
	}, nil
}
