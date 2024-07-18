package infrastructure

import (
	"context"

	"github.com/platform9/cluster-api-sdk-go/infrastructure"
	awsv2 "sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type AWSProviderImpl struct {
	Client client.Client
}

type AWSProvider interface {
	CreateInfraCluster(ctx context.Context, input infrastructure.CreateInfraClusterInput) error
	GetInfraCluster(ctx context.Context, input infrastructure.GetInfraClusterInput) (any, error)
	DeleteInfraCluster(ctx context.Context, input infrastructure.DeleteInfraClusterInput) error

	CreateInfraStaticIdentity(ctx context.Context, input infrastructure.CreateInfraClusterIdentityInput) error
	DeleteInfraStaticIdentity(ctx context.Context, input infrastructure.DeleteInfraClusterIdentityInput) error
	CreateSecretForAWSSI(ctx context.Context, input infrastructure.CreateInfraClusterIdentityInput) error

	DeleteSecretForAWSSI(ctx context.Context, input infrastructure.CreateInfraClusterIdentityInput) error
	CreateClusterRoleIdentity(ctx context.Context, input infrastructure.CreateInfraClusterIdentityInput) error
	DeleteClusterRoleIdentity(ctx context.Context, input infrastructure.DeleteInfraClusterIdentityInput) error
	GetClusterRoleIdentity(ctx context.Context, input infrastructure.GetInfraClusterIdentityInput) (*awsv2.AWSClusterRoleIdentity, error)

	CreateInfraMachine(ctx context.Context, input infrastructure.CreateInfraMachineInput) error
	DeleteInfraMachine(ctx context.Context, input infrastructure.DeleteInfraMachineInput) error
	GetInfraMachine(ctx context.Context, input infrastructure.GetInfraMachineInput) error

	CreateInfraMachineTemplate(ctx context.Context, input infrastructure.CreateInfraMachineTemplateInput) error
	DeleteInfraMachineTemplate(ctx context.Context, input infrastructure.DeleteInfraMachineTemplateInput) error
	GetInfraMachineTemplate(ctx context.Context, input infrastructure.GetInfraMachineTemplateInput) error
}

var _ AWSProvider = (*AWSProviderImpl)(nil)
