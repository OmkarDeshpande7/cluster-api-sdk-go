package infrastructure

import "context"

type InfraProvider interface {
	CreateInfraCluster(ctx context.Context, input CreateInfraClusterInput) error
	CreateInfraStaticIdentity(ctx context.Context, input CreateInfraClusterStaticIdentityInput) error
	CreateInfraMachine(ctx context.Context, input CreateInfraMachineInput) error
	CreateInfraMachineTemplate(ctx context.Context, input CreateInfraMachineTemplateInput) error

	DeleteInfraCluster(ctx context.Context, input DeleteInfraClusterInput) error
	DeleteInfraStaticIdentity(ctx context.Context, input DeleteInfraClusterStaticIdentityInput) error
	DeleteInfraMachine(ctx context.Context, input DeleteInfraMachineInput) error
	DeleteInfraMachineTemplate(ctx context.Context, input DeleteInfraMachineTemplateInput) error
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
