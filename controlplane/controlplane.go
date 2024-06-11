package controlplane

import (
	"context"
)

type ControlPlaneProvider interface {
	GetControlPlane(ctx context.Context, input GetControlPlaneInput) error

	CreateControlPlane(ctx context.Context, input CreateControlPlaneInput) error

	DeleteControlPlane(ctx context.Context, input DeleteControlPlaneInput) error
}
type GetControlPlaneInput interface {
	GetName() string
}

type CreateControlPlaneInput interface {
	GetName() string
}

type DeleteControlPlaneInput interface {
	GetName() string
}
