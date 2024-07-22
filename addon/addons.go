package addon

import (
	"context"
)

type AddonProvider interface {
	GetAddonClusterSummary(ctx context.Context, input GetAddonClusterSummaryInput) (any, error)
}

type GetAddonClusterSummaryInput interface {
	GetName() string
}
