package certs

import "context"

type Interface interface {
	// SearchAppOperator searches for secrets containing TLS certs
	// for managed catalogue service.
	SearchAppOperator(ctx context.Context, clusterID string) (AppOperator, error)
	// SearchCluster searches for secrets containing TLS certs for guest
	// clusters components.
	SearchCluster(ctx context.Context, clusterID string) (Cluster, error)
	// SearchClusterOperator searches for secrets containing TLS certs for
	// connecting to guest clusters.
	SearchClusterOperator(ctx context.Context, clusterID string) (ClusterOperator, error)
	// SearchDraining searches for secrets containing TLS certs for
	// draining nodes in guest clusters.
	SearchDraining(ctx context.Context, clusterID string) (Draining, error)
	// SearchMonitoring searches for secrets containing TLS certs for
	// monitoring guest clusters.
	SearchMonitoring(ctx context.Context, clusterID string) (Monitoring, error)
	// SearchTLS provides a dedicated way to lookup a single TLS asset for one
	// specific purpose. This might be used for e.g. granting guest cluster
	// access within operators.
	SearchTLS(ctx context.Context, clusterID string, cert Cert) (TLS, error)
}
