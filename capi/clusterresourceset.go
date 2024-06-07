package capi

import (
	"context"

	addonsv1 "sigs.k8s.io/cluster-api/exp/addons/api/v1beta1"
)

type CreateClusterResourceSetInput struct {
}

func (c *CAPICore) CreateClusterResourceSet(ctx context.Context, input *CreateClusterResourceSetInput) error {
	crs := &addonsv1.ClusterResourceSet{}
	err := c.Client.Create(ctx, crs)
	if err != nil {
		return err
	}
	return nil
}
