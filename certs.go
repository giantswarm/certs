package certs

// Cert is a certificate name.
type Cert string

func (c Cert) String() string {
	return string(c)
}

// These constants used as Cert parsing a secret received from the API.
const (
	APICert                Cert = "api"
	AppOperatorAPICert     Cert = "app-operator-api"
	AWSOperatorAPICert     Cert = "aws-operator-api"
	CalicoEtcdClientCert   Cert = "calico-etcd-client"
	ClusterOperatorAPICert Cert = "cluster-operator-api"
	EtcdCert               Cert = "etcd"
	Etcd1Cert              Cert = "etcd1"
	Etcd2Cert              Cert = "etcd2"
	Etcd3Cert              Cert = "etcd3"
	FlanneldEtcdClientCert Cert = "flanneld-etcd-client"
	InternalAPICert        Cert = "internal-api"
	NodeOperatorCert       Cert = "node-operator"
	PrometheusCert         Cert = "prometheus"
	ServiceAccountCert     Cert = "service-account"
	WorkerCert             Cert = "worker"
)

// AllCerts lists all certificates that can be created by cert-operator.
var AllCerts = []Cert{
	APICert,
	AppOperatorAPICert,
	AWSOperatorAPICert,
	CalicoEtcdClientCert,
	ClusterOperatorAPICert,
	EtcdCert,
	Etcd1Cert,
	Etcd2Cert,
	Etcd3Cert,
	FlanneldEtcdClientCert,
	InternalAPICert,
	NodeOperatorCert,
	PrometheusCert,
	ServiceAccountCert,
	WorkerCert,
}
