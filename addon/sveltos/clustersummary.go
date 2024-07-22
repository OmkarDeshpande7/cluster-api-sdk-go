package addon

import (
	"context"
	"fmt"

	addon "github.com/platform9/cluster-api-sdk-go/addon"
	sveltos "github.com/projectsveltos/addon-controller/api/v1alpha1"
)

type GetSveltosClusterSummaryInput struct {
	Name, Namespace string
}

func (c GetSveltosClusterSummaryInput) GetName() string {
	return c.Name
}

func (c SveltosProviderImpl) GetAddonClusterSummary(ctx context.Context, input addon.GetAddonClusterSummaryInput) (any, error) {
	sInput, ok := input.(GetSveltosClusterSummaryInput)
	if !ok {
		return nil, fmt.Errorf("invalid argument to GetAddonClusterSummary, input is not type '%s'", "GetSveltosClusterSummaryInput")
	}
	return sveltos.GetClusterSummary(ctx, c.Client, sInput.Namespace, sInput.Name)
}
