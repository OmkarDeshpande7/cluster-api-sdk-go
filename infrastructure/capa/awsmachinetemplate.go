package infrastructure

import (
	"context"

	"github.com/OmkarDeshpande7/cluster-api-sdk-go/infrastructure"
)

type CreateAWSMachineTemplateInput struct {
	Name string
}

func (c *CreateAWSMachineTemplateInput) GetName() string {
	return c.Name
}

func (a *AWSProvider) CreateInfraMachineTemplate(ctx context.Context, input *infrastructure.CreateInfraMachineTemplateInput) error {
	return nil
}
