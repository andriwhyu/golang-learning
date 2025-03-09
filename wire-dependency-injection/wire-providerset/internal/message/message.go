package message

import (
	"fmt"
	"github.com/google/wire"
	"wire-providerset/internal/greeter"
)

var ProviderSet = wire.NewSet(NewMessage)

type Message struct {
	Greeter *greeter.Greeter
}

func NewMessage(g *greeter.Greeter) *Message {
	return &Message{Greeter: g}
}

func (m *Message) Print(name string) {
	fmt.Println(m.Greeter.Greet(name))
}
