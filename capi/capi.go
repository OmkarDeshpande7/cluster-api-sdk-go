package capi

import (
	"sigs.k8s.io/cluster-api/cmd/clusterctl/client"
)

type CAPICore struct {
	client client.Client
}
