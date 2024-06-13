package bootstrap

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type KubeadmProvider struct {
	Client client.Client
}
