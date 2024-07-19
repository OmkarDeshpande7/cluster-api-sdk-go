package bootstrap

import (
	"context"
	"fmt"

	"github.com/platform9/cluster-api-sdk-go/bootstrap"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	bootstrapv1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1beta1"
)

type GetKubeadmConfigTemplateInput struct {
	Name, Namespace string
}

type CreateKubeadmConfigTemplateInput struct {
	Name, Namespace   string
	Files             []bootstrapv1.File
	JoinConfiguration *bootstrapv1.JoinConfiguration
}

type DeleteKubeadmConfigTemplateInput struct {
	Name, Namespace string
}

func (c GetKubeadmConfigTemplateInput) GetName() string {
	return c.Name
}

func (c CreateKubeadmConfigTemplateInput) GetName() string {
	return c.Name
}

func (c DeleteKubeadmConfigTemplateInput) GetName() string {
	return c.Name
}

func (c KubeadmProviderImpl) GetBootstrapConfigTemplate(ctx context.Context, input bootstrap.GetBootstrapConfigTemplateInput) error {
	return nil
}

func (c KubeadmProviderImpl) CreateBootstrapConfigTemplate(ctx context.Context, input bootstrap.CreateBootstrapConfigTemplateInput) error {
	kctInput, ok := input.(CreateKubeadmConfigTemplateInput)
	if !ok {
		return fmt.Errorf("invalid argument to CreateBootstrapConfigTemplate, input is not type '%s'", TypeCreateKubeadmConfigTemplateInput)
	}

	kcp := &bootstrapv1.KubeadmConfigTemplate{
		ObjectMeta: metav1.ObjectMeta{
			Name:      kctInput.Name,
			Namespace: kctInput.Namespace,
		},
		Spec: bootstrapv1.KubeadmConfigTemplateSpec{
			Template: bootstrapv1.KubeadmConfigTemplateResource{
				Spec: bootstrapv1.KubeadmConfigSpec{},
			},
		},
	}

	kcp.Spec.Template.Spec.Files = kctInput.Files
	kcp.Spec.Template.Spec.JoinConfiguration = kctInput.JoinConfiguration
	err := c.Client.Create(ctx, kcp)
	if err != nil {
		return err
	}
	return nil
}

func (c KubeadmProviderImpl) DeleteBootstrapConfigTemplate(ctx context.Context, input bootstrap.DeleteBootstrapConfigTemplateInput) error {
	return nil
}
