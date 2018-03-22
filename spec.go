package certs

type Interface interface {
	// SearchCluster searches for secrets containing TLS certs for guest
	// clusters components.
	SearchCluster(clusterID string) (Cluster, error)
	// SearchGuestResources searches for secrets containing TLS certs for
	// managing resources in guest clusters.
	SearchGuestResources(clusterID string) (GuestResources, error)
	// SearchDraining searches for secrets containing TLS certs for
	// draining nodes in guest clusters.
	SearchDraining(clusterID string) (Draining, error)
	// SearchMonitoring searches for secrets containing TLS certs for
	// monitoring guest clusters.
	SearchMonitoring(clusterID string) (Monitoring, error)
}
