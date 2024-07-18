package controlplane

import (
	"context"

	"github.com/platform9/cluster-api-sdk-go/controlplane"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type KamajiProviderImpl struct {
	Client client.Client
}

type KamajiProvider interface {
	GetControlPlane(ctx context.Context, input controlplane.GetControlPlaneInput) (any, error)
	CreateControlPlane(ctx context.Context, input controlplane.CreateControlPlaneInput) error
	DeleteControlPlane(ctx context.Context, input controlplane.DeleteControlPlaneInput) error
}

var _ KamajiProvider = (*KamajiProviderImpl)(nil)
