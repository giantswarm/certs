package certs

import (
	"fmt"
	"time"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

const (
	// DefaultWatchTimeOut is the time to wait on watches against the Kubernetes
	// API before giving up and throwing an error.
	DefaultWatchTimeOut = 90 * time.Second
)

type Config struct {
	K8sClient kubernetes.Interface
	Logger    micrologger.Logger

	WatchTimeOut time.Duration
}

type Searcher struct {
	k8sClient kubernetes.Interface
	logger    micrologger.Logger

	watchTimeOut time.Duration
}

func NewSearcher(config Config) (*Searcher, error) {
	if config.K8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.K8sClient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.WatchTimeOut == 0 {
		config.WatchTimeOut = DefaultWatchTimeOut
	}

	s := &Searcher{
		k8sClient: config.K8sClient,
		logger:    config.Logger,

		watchTimeOut: config.WatchTimeOut,
	}

	return s, nil
}

func (s *Searcher) SearchCluster(clusterID string) (Cluster, error) {
	var cluster Cluster

	certificates := []struct {
		TLS  *TLS
		Cert Cert
	}{
		{TLS: &cluster.APIServer, Cert: APICert},
		{TLS: &cluster.CalicoClient, Cert: CalicoCert},
		{TLS: &cluster.EtcdServer, Cert: EtcdCert},
		{TLS: &cluster.ServiceAccount, Cert: ServiceAccountCert},
		{TLS: &cluster.Worker, Cert: WorkerCert},
	}

	for _, c := range certificates {
		err := s.search(c.TLS, clusterID, c.Cert)
		if err != nil {
			return Cluster{}, microerror.Mask(err)
		}
	}

	return cluster, nil
}

func (s *Searcher) SearchClusterOperator(clusterID string) (ClusterOperator, error) {
	var clusterOperator ClusterOperator

	certificates := []struct {
		TLS  *TLS
		Cert Cert
	}{
		{TLS: &clusterOperator.APIServer, Cert: ClusterOperatorAPICert},
	}

	for _, c := range certificates {
		err := s.search(c.TLS, clusterID, c.Cert)
		if err != nil {
			return ClusterOperator{}, microerror.Mask(err)
		}
	}

	return clusterOperator, nil
}

func (s *Searcher) SearchDraining(clusterID string) (Draining, error) {
	var draining Draining

	certificates := []struct {
		TLS  *TLS
		Cert Cert
	}{
		{TLS: &draining.NodeOperator, Cert: NodeOperatorCert},
	}

	for _, c := range certificates {
		err := s.search(c.TLS, clusterID, c.Cert)
		if err != nil {
			return Draining{}, microerror.Mask(err)
		}
	}

	return draining, nil
}

func (s *Searcher) SearchMonitoring(clusterID string) (Monitoring, error) {
	var monitoring Monitoring

	certificates := []struct {
		TLS  *TLS
		Cert Cert
	}{
		{TLS: &monitoring.Prometheus, Cert: PrometheusCert},
	}

	for _, c := range certificates {
		err := s.search(c.TLS, clusterID, c.Cert)
		if err != nil {
			return Monitoring{}, microerror.Mask(err)
		}
	}

	return monitoring, nil
}

func (s *Searcher) searchError(tls *TLS, clusterID string, cert Cert, err error) error {
	if err != nil {
		return err
	}
	return s.search(tls, clusterID, cert)
}

func (s *Searcher) search(tls *TLS, clusterID string, cert Cert) error {
	// Select only secrets that match the given certificate and the given
	// cluster clusterID.
	selector := fmt.Sprintf("%s=%s, %s=%s", certficateLabel, cert, clusterIDLabel, clusterID)

	watcher, err := s.k8sClient.Core().Secrets(SecretNamespace).Watch(metav1.ListOptions{
		LabelSelector: selector,
	})
	if err != nil {
		return microerror.Mask(err)
	}

	defer watcher.Stop()

	for {
		select {
		case event, ok := <-watcher.ResultChan():
			if !ok {
				return microerror.Maskf(executionError, "watching secrets, selector = %q: unexpected closed channel", selector)
			}

			switch event.Type {
			case watch.Added:
				err := fillTLSFromSecret(tls, event.Object, clusterID, cert)
				if err != nil {
					return microerror.Maskf(err, "watching secrets, selector = %q")
				}

				return nil
			case watch.Deleted:
				// Noop. Ignore deleted events. These are
				// handled by the certificate operator.
			case watch.Error:
				return microerror.Maskf(executionError, "watching secrets, selector = %q: %v", selector, apierrors.FromObject(event.Object))
			}
		case <-time.After(s.watchTimeOut):
			return microerror.Maskf(timeoutError, "waiting secrets, selector = %q", selector)
		}
	}
}

func fillTLSFromSecret(tls *TLS, obj runtime.Object, clusterID string, cert Cert) error {
	secret, ok := obj.(*corev1.Secret)
	if !ok || secret == nil {
		return microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", secret, obj)
	}

	gotClusterID := secret.Labels[clusterIDLabel]
	if clusterID != gotClusterID {
		return microerror.Maskf(invalidSecretError, "expected clusterID = %q, got %q", clusterID, gotClusterID)
	}
	gotcert := secret.Labels[certficateLabel]
	if string(cert) != gotcert {
		return microerror.Maskf(invalidSecretError, "expected certificate = %q, got %q", cert, gotcert)
	}

	if tls.CA, ok = secret.Data["ca"]; !ok {
		return microerror.Maskf(invalidSecretError, "%q key missing", "ca")
	}
	if tls.Crt, ok = secret.Data["crt"]; !ok {
		return microerror.Maskf(invalidSecretError, "%q key missing", "crt")
	}
	if tls.Key, ok = secret.Data["key"]; !ok {
		return microerror.Maskf(invalidSecretError, "%q key missing", "key")
	}

	return nil
}
