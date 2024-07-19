package bootstrap

import (
	"context"

	"github.com/platform9/cluster-api-sdk-go/bootstrap"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type KubeadmProviderImpl struct {
	Client client.Client
}

type KubeadmProvider interface {
	GetBootstrapConfigTemplate(ctx context.Context, input bootstrap.GetBootstrapConfigTemplateInput) error
	CreateBootstrapConfigTemplate(ctx context.Context, input bootstrap.CreateBootstrapConfigTemplateInput) error
	DeleteBootstrapConfigTemplate(ctx context.Context, input bootstrap.DeleteBootstrapConfigTemplateInput) error
}

var _ KubeadmProvider = (*KubeadmProviderImpl)(nil)
