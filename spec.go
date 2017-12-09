package certificates

type Interface interface {
	SearchCluster(clusterID string) (Cluster, error)
	SearchMonitoring(clusterID string) (Monitoring, error)
}
