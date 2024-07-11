//go:generate go run -mod=mod go.uber.org/mock/mockgen -package mocks -destination=./mock_core.go -source=../../capi/capi.go -build_flags=-mod=mod
//go:generate go run -mod=mod go.uber.org/mock/mockgen -package mocks -destination=./mock_controlplane.go -source=../../controlplane/controlplane.go -build_flags=-mod=mod
//go:generate go run -mod=mod go.uber.org/mock/mockgen -package mocks -destination=./mock_infrastructure.go -source=../../infrastructure/infrastructure.go -build_flags=-mod=mod

package mocks

import _ "go.uber.org/mock/mockgen/model"
