package certs

type File struct {
	AbsolutePath string
	Data         []byte
}

func NewFilesAPI(cert TLS) []File {
	return []File{
		// Kubernetes API server.
		{
			AbsolutePath: "/etc/kubernetes/ssl/apiserver-ca.pem",
			Data:         cert.CA,
		},
		{
			AbsolutePath: "/etc/kubernetes/ssl/apiserver-crt.pem",
			Data:         cert.Crt,
		},
		{
			AbsolutePath: "/etc/kubernetes/ssl/apiserver-key.pem",
			Data:         cert.Key,
		},
	}
}

func NewFilesCalicoEtcdClient(cert TLS) []File {
	return []File{
		// Calico Etcd client.
		{
			AbsolutePath: "/etc/kubernetes/ssl/calico/etcd-ca",
			Data:         cert.CA,
		},
		{
			AbsolutePath: "/etc/kubernetes/ssl/calico/etcd-cert",
			Data:         cert.Crt,
		},
		{
			AbsolutePath: "/etc/kubernetes/ssl/calico/etcd-key",
			Data:         cert.Key,
		},
	}
}

func NewFilesEtcd(cert TLS) []File {
	return []File{
		// Etcd server.
		{
			AbsolutePath: "/etc/kubernetes/ssl/etcd/server-ca.pem",
			Data:         cert.CA,
		},
		{
			AbsolutePath: "/etc/kubernetes/ssl/etcd/server-crt.pem",
			Data:         cert.Crt,
		},
		{
			AbsolutePath: "/etc/kubernetes/ssl/etcd/server-key.pem",
			Data:         cert.Key,
		},
	}
}

func NewFilesPrometheusEtcdClient(cert TLS) []File {
	return []File{
		// Prometheus Etcd client.
		{
			AbsolutePath: "/etc/kubernetes/ssl/etcd/client-ca.pem",
			Data:         cert.CA,
		},
		{
			AbsolutePath: "/etc/kubernetes/ssl/etcd/client-crt.pem",
			Data:         cert.Crt,
		},
		{
			AbsolutePath: "/etc/kubernetes/ssl/etcd/client-key.pem",
			Data:         cert.Key,
		},
	}
}

func NewFilesServiceAccount(cert TLS) []File {
	return []File{
		// Service account (only key file is used).
		{
			AbsolutePath: "/etc/kubernetes/ssl/service-account-key.pem",
			Data:         cert.Key,
		},
	}
}

func NewFilesWorker(cert TLS) []File {
	return []File{
		{
			Data:         cert.CA,
			AbsolutePath: "/etc/kubernetes/ssl/worker-ca.pem",
		},
		{
			Data:         cert.Crt,
			AbsolutePath: "/etc/kubernetes/ssl/worker-crt.pem",
		},
		{
			Data:         cert.Key,
			AbsolutePath: "/etc/kubernetes/ssl/worker-key.pem",
		},
	}
}
