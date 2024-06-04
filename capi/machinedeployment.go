package capi

import "context"

type CreateMachineDeploymentInput struct {
	Name, Namespace string
	// Count of minimum  machines/instances required in pool
	MinMachines int64 `json:"minMachines"`

	// Count of maximum machines/instances allowed to create in pool
	MaxMachines int64 `json:"maxMachines"`
}

func (c *CAPICore) CreateMachineDeployment(ctx context.Context, input *CreateMachineDeploymentInput) error {
	return nil
}
