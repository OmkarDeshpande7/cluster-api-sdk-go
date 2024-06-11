package capi

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

type CreateClusterInput struct {
	Name, Namespace, ControlPlaneFQDN  string
	ControlPlaneHostPort               int32
	PodCIDRs, ServiceCIDRs             []string
	ControlPlaneRef, InfrastructureRef *corev1.ObjectReference
}

type GetClusterInput struct {
	Name, Namespace string
}

type DeleteClusterInput struct {
	Name, Namespace string
}

func (c *CAPICore) CreateCluster(ctx context.Context, input *CreateClusterInput) error {
	cluster := &clusterv1.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      input.Name,
			Namespace: input.Namespace,
		},
		Spec: clusterv1.ClusterSpec{
			ControlPlaneEndpoint: clusterv1.APIEndpoint{
				Host: input.ControlPlaneFQDN,
				Port: input.ControlPlaneHostPort,
			},
			ClusterNetwork: &clusterv1.ClusterNetwork{
				APIServerPort: &input.ControlPlaneHostPort,
			},
			ControlPlaneRef:   input.ControlPlaneRef,
			InfrastructureRef: input.InfrastructureRef,
		},
	}
	if len(input.PodCIDRs) > 0 {
		cluster.Spec.ClusterNetwork.Pods = &clusterv1.NetworkRanges{
			CIDRBlocks: input.PodCIDRs,
		}
	}

	if len(input.ServiceCIDRs) > 0 {
		cluster.Spec.ClusterNetwork.Services = &clusterv1.NetworkRanges{
			CIDRBlocks: input.ServiceCIDRs,
		}
	}

	err := c.Client.Create(ctx, cluster)
	if err != nil {
		return err
	}
	return nil
}

func (c *CAPICore) GetCluster(ctx context.Context, input *GetClusterInput) (*clusterv1.Cluster, error) {
	cluster := &clusterv1.Cluster{}
	err := c.Client.Get(ctx, types.NamespacedName{
		Name:      input.Name,
		Namespace: input.Namespace,
	}, cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

func (c *CAPICore) DeleteCluster(ctx context.Context, input *DeleteClusterInput) error {
	cluster := &clusterv1.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      input.Name,
			Namespace: input.Namespace,
		},
	}

	err := c.Client.Delete(ctx, cluster)
	if err != nil {
		return err
	}
	return nil
}
