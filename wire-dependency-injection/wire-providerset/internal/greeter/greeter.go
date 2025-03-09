package greeter

import (
	"fmt"
	"github.com/google/wire"
	"wire-providerset/internal/config"
)

var ProviderSet = wire.NewSet(NewGreeter)

type Greeter struct {
	Config *config.Config
}

func NewGreeter(c *config.Config) *Greeter {
	return &Greeter{Config: c}
}

func (g *Greeter) Greet(name string) string {
	return fmt.Sprintf("%s, %s!", g.Config.Greating, name)
}
