package domain

// Karmada 抽象类，用于与Karmada集群交互
type Karmada interface {
	GetClusters() ([]string, error)
	GetResources() (map[string]interface{}, error)
	GetResourceSummary() (map[string]interface{}, error)
	GetDeploymentsByModelTag(modelTag string) ([]interface{}, error)
}
