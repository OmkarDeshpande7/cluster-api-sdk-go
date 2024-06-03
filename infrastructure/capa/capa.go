package infrastructure

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type AWSProvider struct {
	Client client.Client
}
