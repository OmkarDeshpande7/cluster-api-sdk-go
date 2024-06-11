package capi

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

type GetClusterInput struct {
	Name, Namespace string
}

type CreateClusterInput struct {
	Name, Namespace                    string
	ControlPlaneFQDN                   *string
	APIServerPort                      *int32
	PodCIDRs, ServiceCIDRs             []string
	ControlPlaneRef, InfrastructureRef *corev1.ObjectReference
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

			ClusterNetwork: &clusterv1.ClusterNetwork{
				APIServerPort: input.APIServerPort,
			},
			ControlPlaneRef:   input.ControlPlaneRef,
			InfrastructureRef: input.InfrastructureRef,
		},
	}

	if input.ControlPlaneFQDN != nil {
		cluster.Spec.ControlPlaneEndpoint.Host = *input.ControlPlaneFQDN
	}
	if input.APIServerPort != nil {
		cluster.Spec.ControlPlaneEndpoint.Port = *input.APIServerPort
		if cluster.Spec.ClusterNetwork == nil {
			cluster.Spec.ClusterNetwork = &clusterv1.ClusterNetwork{
				APIServerPort: input.APIServerPort,
			}
		} else {
			cluster.Spec.ClusterNetwork.APIServerPort = input.APIServerPort
		}
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
