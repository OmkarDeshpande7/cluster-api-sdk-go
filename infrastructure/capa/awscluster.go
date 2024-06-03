package infrastructure

import (
	"context"

	"github.com/OmkarDeshpande7/cluster-api-sdk-go/infrastructure"
)

type CreateAWSClusterInput struct {
	Name string
}

func (c *CreateAWSClusterInput) GetName() string {
	return c.Name
}

func (a *AWSProvider) CreateInfraCluster(ctx context.Context, input *infrastructure.CreateInfraClusterInput) error {
	return nil
}
