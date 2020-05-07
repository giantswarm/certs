package certs

type TLS struct {
	CA, Crt, Key []byte
}

type AppOperator struct {
	APIServer TLS
}

type ClusterOperator struct {
	APIServer TLS
}

type Draining struct {
	NodeOperator TLS
}

type Monitoring struct {
	KubeStateMetrics TLS
	Prometheus       TLS
}
