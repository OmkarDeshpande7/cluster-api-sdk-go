package infrastructure

import (
	"context"

	awsv2 "sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"

	"github.com/OmkarDeshpande7/cluster-api-sdk-go/infrastructure"
)

type CreateAWSMachineInput struct {
	Name, Namespace string
}

func (c *CreateAWSMachineInput) GetName() string {
	return c.Name
}

func (a *AWSProvider) CreateInfraMachine(ctx context.Context, input infrastructure.CreateInfraMachineInput) error {
	awsMachine := &awsv2.AWSMachine{}
	err := a.Client.Create(ctx, awsMachine)
	if err != nil {
		return err
	}
	return nil
}
