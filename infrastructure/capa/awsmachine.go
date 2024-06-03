package infrastructure

import (
	"context"

	"github.com/OmkarDeshpande7/cluster-api-sdk-go/infrastructure"
)

type CreateAWSMachineInput struct {
	Name string
}

func (c *CreateAWSMachineInput) GetName() string {
	return c.Name
}

func (a *AWSProvider) CreateInfraMachine(ctx context.Context, input *infrastructure.CreateInfraMachineInput) error {
	return nil
}
