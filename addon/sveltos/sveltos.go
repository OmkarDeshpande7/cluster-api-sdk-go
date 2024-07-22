package addon

import (
	"context"

	addon "github.com/platform9/cluster-api-sdk-go/addon"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SveltosProviderImpl struct {
	Client client.Client
}

type SveltosProvider interface {
	GetAddonClusterSummary(ctx context.Context, input addon.GetAddonClusterSummaryInput) (any, error)
}

var _ SveltosProvider = (*SveltosProviderImpl)(nil)
