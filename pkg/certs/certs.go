package certs

// Cert refers to a component we generate a certificate for.
type Cert string

func (c Cert) String() string {
	return string(c)
}

// These constants are used as components identifying the purpose of generated
// certificates.
const (
	APICert                  Cert = "api"
	AppOperatorAPICert       Cert = "app-operator-api"
	AWSOperatorAPICert       Cert = "aws-operator-api"
	CalicoEtcdClientCert     Cert = "calico-etcd-client"
	ClusterOperatorAPICert   Cert = "cluster-operator-api"
	EtcdCert                 Cert = "etcd"
	Etcd1Cert                Cert = "etcd1"
	Etcd2Cert                Cert = "etcd2"
	Etcd3Cert                Cert = "etcd3"
	FlanneldEtcdClientCert   Cert = "flanneld-etcd-client"
	InternalAPICert          Cert = "internal-api"
	NodeOperatorCert         Cert = "node-operator"
	PrometheusCert           Cert = "prometheus"
	PrometheusEtcdClientCert Cert = "prometheus-etcd-client"
	ServiceAccountCert       Cert = "service-account"
	WorkerCert               Cert = "worker"
)
