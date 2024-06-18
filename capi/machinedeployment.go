package capi

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/utils/ptr"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CreateMachineDeploymentInput struct {
	Name, Namespace, ClusterName string
	// Count of minimum  machines/instances required in pool
	MinMachines int32 `json:"minMachines"`

	// Count of maximum machines/instances allowed to create in pool
	MaxMachines int32 `json:"maxMachines"`

	BootstrapRef *corev1.ObjectReference `json:"bootstrapRef"`

	InfrastructureRef *corev1.ObjectReference `json:"infrastructureRef"`

	NodeVersion *string `json:"nodeVersion"`

	MatchLabels map[string]string
}

type ListMachineDeploymentInput struct {
	Namespace     string
	LabelSelector labels.Selector
}

func (c *CAPICore) CreateMachineDeployment(ctx context.Context, input *CreateMachineDeploymentInput) error {
	md := clusterv1.MachineDeployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      input.Name,
			Namespace: input.Namespace,
		},
		Spec: clusterv1.MachineDeploymentSpec{
			Selector: metav1.LabelSelector{
				MatchLabels: input.MatchLabels,
			},
			Replicas:    ptr.To[int32](input.MinMachines),
			ClusterName: input.ClusterName,
			Template: clusterv1.MachineTemplateSpec{
				Spec: clusterv1.MachineSpec{
					Bootstrap: clusterv1.Bootstrap{
						ConfigRef: input.BootstrapRef,
					},
					InfrastructureRef: *input.InfrastructureRef,
					ClusterName:       input.ClusterName,
					Version:           input.NodeVersion,
				},
			},
		},
	}
	err := c.Client.Create(ctx, &md)
	if err != nil {
		return err
	}
	return nil
}

func (c *CAPICore) ListMachineDeployment(ctx context.Context, input *ListMachineDeploymentInput) (*clusterv1.MachineDeploymentList, error) {
	mdList := &clusterv1.MachineDeploymentList{}
	err := c.Client.List(ctx, mdList, &client.ListOptions{
		Namespace:     input.Namespace,
		LabelSelector: input.LabelSelector,
	})
	if err != nil {
		return nil, err
	}
	return mdList, nil
}
