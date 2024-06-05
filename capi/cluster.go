package capi

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

type CreateClusterInput struct {
	Name, Namespace                    string
	PodCIDRs, ServiceCIDRs             []string
	ControlPlaneRef, InfrastructureRef *corev1.ObjectReference
}

func (c *CAPICore) CreateCluster(ctx context.Context, input *CreateClusterInput) error {
	cluster := &clusterv1.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      input.Name,
			Namespace: input.Namespace,
		},
		Spec: clusterv1.ClusterSpec{
			ClusterNetwork: &clusterv1.ClusterNetwork{
				Pods: &clusterv1.NetworkRanges{
					CIDRBlocks: input.PodCIDRs,
				},
				Services: &clusterv1.NetworkRanges{
					CIDRBlocks: input.ServiceCIDRs,
				},
			},
			ControlPlaneRef:   input.ControlPlaneRef,
			InfrastructureRef: input.InfrastructureRef,
		},
	}
	err := c.Client.Create(ctx, cluster)
	if err != nil {
		return err
	}
	return nil
}
