package capi

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CAPICore struct {
	Client client.Client
}
