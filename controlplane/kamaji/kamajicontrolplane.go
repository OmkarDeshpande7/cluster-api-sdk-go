package controlplane

import (
	"context"
	"fmt"

	kamajiv1alpha1 "github.com/clastix/kamaji/api/v1alpha1"
	"github.com/platform9/cluster-api-sdk-go/controlplane"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/pointer"

	kcpv1alpha1 "github.com/clastix/cluster-api-control-plane-provider-kamaji/api/v1alpha1"
)

type GetKamajiControlPlaneInput struct {
	Name, Namespace string
}

type CreateKamajiControlPlaneInput struct {
	Name, Namespace, Datastore, CGroupDriver, K8sVersion, IngressHostname string
	APIServerExtraArgs, ControllerManagerExtraArgs, CertSANs              []string
	AddonsSpec                                                            *kamajiv1alpha1.AddonsSpec
	CoreDNSAddonSpec                                                      *kcpv1alpha1.CoreDNSAddonSpec
	Replicas                                                              int32
	ExtraAnnotations                                                      map[string]string
}

type DeleteKamajiControlPlaneInput struct {
	Name, Namespace string
}

func (c GetKamajiControlPlaneInput) GetName() string {
	return c.Name
}

func (c CreateKamajiControlPlaneInput) GetName() string {
	return c.Name
}

func (c DeleteKamajiControlPlaneInput) GetName() string {
	return c.Name
}

func (c *KamajiProvider) GetControlPlane(ctx context.Context, input controlplane.GetControlPlaneInput) (any, error) {
	cpInput, ok := input.(GetKamajiControlPlaneInput)
	if !ok {
		return nil, fmt.Errorf("invalid argument to GetControlPlane, input is not type '%s'", TypeGetKamajiControlPlaneInput)
	}
	controlPlane := &kcpv1alpha1.KamajiControlPlane{}
	err := c.Client.Get(ctx, types.NamespacedName{
		Name:      cpInput.Name,
		Namespace: cpInput.Namespace,
	}, controlPlane)
	if err != nil {
		return nil, err
	}
	return controlPlane, nil
}

func (c *KamajiProvider) CreateControlPlane(ctx context.Context, input controlplane.CreateControlPlaneInput) error {
	cpInput, ok := input.(CreateKamajiControlPlaneInput)
	if !ok {
		return fmt.Errorf("invalid argument to CreateControlPlane, input is not type '%s'", TypeCreateKamajiControlPlaneInput)
	}

	kcp := &kcpv1alpha1.KamajiControlPlane{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cpInput.Name,
			Namespace: cpInput.Namespace,
		},
		Spec: kcpv1alpha1.KamajiControlPlaneSpec{
			ApiServer: kcpv1alpha1.ControlPlaneComponent{
				ExtraArgs: cpInput.APIServerExtraArgs,
			},
			ControllerManager: kcpv1alpha1.ControlPlaneComponent{
				ExtraArgs: cpInput.ControllerManagerExtraArgs,
			},
			DataStoreName: cpInput.Datastore,
			Addons: kcpv1alpha1.AddonsSpec{
				AddonsSpec: *cpInput.AddonsSpec,
				CoreDNS:    cpInput.CoreDNSAddonSpec,
			},
			Kubelet: kamajiv1alpha1.KubeletSpec{
				CGroupFS: kamajiv1alpha1.CGroupDriver(cpInput.CGroupDriver),
			},
			Network: kcpv1alpha1.NetworkComponent{
				ServiceType: kamajiv1alpha1.ServiceTypeClusterIP,
				CertSANs:    cpInput.CertSANs,
				Ingress: &kcpv1alpha1.IngressComponent{
					ClassName:        ClassNameNginx,
					Hostname:         cpInput.IngressHostname,
					ExtraAnnotations: cpInput.ExtraAnnotations,
				},
			},
			Replicas: pointer.Int32(cpInput.Replicas),
			Version:  cpInput.K8sVersion,
		},
	}

	kcp.Spec.Addons.AddonsSpec = *cpInput.AddonsSpec

	err := c.Client.Create(ctx, kcp)
	if err != nil {
		return err
	}
	return nil
}

func (c *KamajiProvider) DeleteControlPlane(ctx context.Context, input controlplane.DeleteControlPlaneInput) error {
	return nil
}
