package infrastructure

import (
	"context"

	"github.com/OmkarDeshpande7/cluster-api-sdk-go/infrastructure"
)

type CreateAWSClusterStaticIdentityInput struct {
	Name string
}

func (c *CreateAWSClusterStaticIdentityInput) GetName() string {
	return c.Name
}
func (a *AWSProvider) CreateInfraStaticIdentity(ctx context.Context, input *infrastructure.CreateInfraClusterStaticIdentityInput) error {
	return nil
}
