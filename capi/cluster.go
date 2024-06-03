package capi

import "context"

type CreateClusterInput struct {
}

func (c *CAPICore) CreateCluster(ctx context.Context, input *CreateClusterInput) error {
	return nil
}
