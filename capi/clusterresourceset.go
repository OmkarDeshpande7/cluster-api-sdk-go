package capi

import (
	"context"

	addonsv1alpha1 "sigs.k8s.io/cluster-api-addon-provider-helm/api/v1alpha1"
)

type CreateClusterResourceSetInput struct {
}

func (c *CAPICore) CreateClusterResourceSet(ctx context.Context, input *CreateClusterResourceSetInput) error {
	crs := addonsv1alpha1.ClusterResourceSet{}
	err := c.Client.Create(ctx, crs)
	if err != nil {
		return err
	}
	return nil
}
