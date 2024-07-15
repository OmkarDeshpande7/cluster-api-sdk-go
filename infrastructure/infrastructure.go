package infrastructure

import (
	"context"
	"errors"

	awsv2 "sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"
)

type InfraProvider interface {
	GetInfraCluster(ctx context.Context, input GetInfraClusterInput) (any, error)
	// GetInfraStaticIdentity(ctx context.Context, input GetInfraClusterStaticIdentityInput) error
	GetInfraMachine(ctx context.Context, input GetInfraMachineInput) error
	GetInfraMachineTemplate(ctx context.Context, input GetInfraMachineTemplateInput) error

	CreateInfraCluster(ctx context.Context, input CreateInfraClusterInput) error
	CreateInfraStaticIdentity(ctx context.Context, input CreateInfraClusterIdentityInput) error
	CreateInfraMachine(ctx context.Context, input CreateInfraMachineInput) error
	CreateInfraMachineTemplate(ctx context.Context, input CreateInfraMachineTemplateInput) error

	DeleteInfraCluster(ctx context.Context, input DeleteInfraClusterInput) error
	DeleteInfraStaticIdentity(ctx context.Context, input DeleteInfraClusterIdentityInput) error
	DeleteInfraMachine(ctx context.Context, input DeleteInfraMachineInput) error
	DeleteInfraMachineTemplate(ctx context.Context, input DeleteInfraMachineTemplateInput) error

	CreateClusterRoleIdentity(ctx context.Context, input CreateInfraClusterIdentityInput) error
	GetClusterRoleIdentity(ctx context.Context, input GetInfraClusterIdentityInput) (*awsv2.AWSClusterRoleIdentity, error)
	DeleteClusterRoleIdentity(ctx context.Context, input DeleteInfraClusterIdentityInput) error
}

var (
	ErrInvalidParameterType = errors.New("input parameter doesn't match the required type")
)

type GetInfraClusterInput interface {
	GetName() string
}
type GetInfraClusterStaticIdentityInput interface {
	GetName() string
}
type GetInfraMachineInput interface {
	GetName() string
}
type GetInfraMachineTemplateInput interface {
	GetName() string
}
type CreateInfraClusterInput interface {
	GetName() string
}
type CreateInfraClusterIdentityInput interface {
	GetName() string
}
type CreateInfraMachineInput interface {
	GetName() string
}
type CreateInfraMachineTemplateInput interface {
	GetName() string
}

type DeleteInfraClusterInput interface {
	GetName() string
}

type DeleteInfraMachineInput interface {
	GetName() string
}
type DeleteInfraMachineTemplateInput interface {
	GetName() string
}

type DeleteInfraClusterIdentityInput interface {
	GetName() string
}

type GetInfraClusterIdentityInput interface {
	GetName() string
}
