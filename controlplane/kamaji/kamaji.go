package controlplane

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type KamajiProvider struct {
	Client client.Client
}
