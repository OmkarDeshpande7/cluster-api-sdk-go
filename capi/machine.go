package capi

import "context"

type CreateMachineInput struct {
}

func (c *CAPICore) CreateMachine(ctx context.Context, input *CreateMachineInput) error {
	return nil
}
