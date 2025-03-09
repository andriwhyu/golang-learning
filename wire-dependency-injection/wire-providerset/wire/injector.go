//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"wire-providerset/internal/config"
	"wire-providerset/internal/greeter"
	"wire-providerset/internal/message"
)

var providerSet = wire.NewSet(
	config.ProviderSet,
	greeter.ProviderSet,
	message.ProviderSet,
)

func InitializeDependency() *message.Message {
	panic(wire.Build(providerSet))
}
