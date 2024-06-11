package infrastructure

import "context"

type InfraProvider interface {
	GetInfraCluster(ctx context.Context, input GetInfraClusterInput) error
	GetInfraStaticIdentity(ctx context.Context, input GetInfraClusterStaticIdentityInput) error
	GetInfraMachine(ctx context.Context, input GetInfraMachineInput) error
	GetInfraMachineTemplate(ctx context.Context, input GetInfraMachineTemplateInput) error

	CreateInfraCluster(ctx context.Context, input CreateInfraClusterInput) error
	CreateInfraStaticIdentity(ctx context.Context, input CreateInfraClusterStaticIdentityInput) error
	CreateInfraMachine(ctx context.Context, input CreateInfraMachineInput) error
	CreateInfraMachineTemplate(ctx context.Context, input CreateInfraMachineTemplateInput) error

	DeleteInfraCluster(ctx context.Context, input DeleteInfraClusterInput) error
	DeleteInfraStaticIdentity(ctx context.Context, input DeleteInfraClusterStaticIdentityInput) error
	DeleteInfraMachine(ctx context.Context, input DeleteInfraMachineInput) error
	DeleteInfraMachineTemplate(ctx context.Context, input DeleteInfraMachineTemplateInput) error
}

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
type CreateInfraClusterStaticIdentityInput interface {
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
type DeleteInfraClusterStaticIdentityInput interface {
	GetName() string
}
type DeleteInfraMachineInput interface {
	GetName() string
}
type DeleteInfraMachineTemplateInput interface {
	GetName() string
}
