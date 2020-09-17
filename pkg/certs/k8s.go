package certs

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// SecretNamespace is the namespace in which secrets are watched.
	SecretNamespace = metav1.NamespaceAll
)

// These constants are used when filtering the secrets, to only retrieve the
// ones we are interested in.
const (
	// certificateLabel is the label used in the secret to identify a secret
	// containing the certificate.
	certificateLabel = "giantswarm.io/certificate"
	// clusterLabel is the label used in the secret to identify a secret
	// containing the certificate.
	clusterLabel = "giantswarm.io/cluster"
)

// K8sName returns Kubernetes object name for the certificate name and
// the guest cluster ID.
func K8sName(cluster string, certificate Cert) string {
	return fmt.Sprintf("%s-%s", cluster, certificate)
}

// K8sLabels returns labels for the Kubernetes  object for the certificate name
// and the guest cluster ID.
func K8sLabels(cluster string, certificate Cert) map[string]string {
	return map[string]string{
		certificateLabel: string(certificate),
		clusterLabel:     cluster,
	}
}
