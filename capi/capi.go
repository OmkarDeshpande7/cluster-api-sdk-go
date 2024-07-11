package capi

import (
	"context"

	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CAPICoreComponents interface {
	CreateCluster(ctx context.Context, input *CreateClusterInput) error
	GetCluster(ctx context.Context, input *GetClusterInput) (*clusterv1.Cluster, error)
	DeleteCluster(ctx context.Context, input *DeleteClusterInput) error

	CreateMachineDeployment(ctx context.Context, input *CreateMachineDeploymentInput) error
	GetMachineDeployment(ctx context.Context, input *GetMachineDeploymentInput) (*clusterv1.MachineDeployment, error)
	UpdateMachineDeployment(ctx context.Context, input *UpdateMachineDeploymentInput) error
	ListMachineDeployment(ctx context.Context, input *ListMachineDeploymentInput) (*clusterv1.MachineDeploymentList, error)

	CreateClusterResourceSet(ctx context.Context, input *CreateClusterResourceSetInput) error
	CreateMachine(ctx context.Context, input *CreateMachineInput) error
	CreateMachinePool(ctx context.Context, input *CreateMachinePoolInput) error
}
type CAPICore struct {
	Client client.Client
}

var _ CAPICoreComponents = (*CAPICore)(nil)
