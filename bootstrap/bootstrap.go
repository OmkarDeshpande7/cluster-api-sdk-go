package bootstrap

import (
	"context"
)

type BootstrapProvider interface {
	GetBootstrapConfigTemplate(ctx context.Context, input GetBootstrapConfigTemplateInput) error

	CreateBootstrapConfigTemplate(ctx context.Context, input CreateBootstrapConfigTemplateInput) error

	DeleteBootstrapConfigTemplate(ctx context.Context, input DeleteBootstrapConfigTemplateInput) error
}
type GetBootstrapConfigTemplateInput interface {
	GetName() string
}

type CreateBootstrapConfigTemplateInput interface {
	GetName() string
}

type DeleteBootstrapConfigTemplateInput interface {
	GetName() string
}
