//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
)

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}