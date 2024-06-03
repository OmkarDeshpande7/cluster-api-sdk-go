package infrastructure

import "context"

type InfraProvider interface {
	CreateInfraCluster(ctx context.Context, input *CreateInfraClusterInput) error
	CreateInfraStaticIdentity(ctx context.Context, input *CreateInfraClusterStaticIdentityInput) error
	CreateInfraMachine(ctx context.Context, input *CreateInfraMachineInput) error
	CreateInfraMachineTemplate(ctx context.Context, input *CreateInfraMachineTemplateInput) error
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
