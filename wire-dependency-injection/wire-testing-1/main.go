package main

import "log"

type Message struct {
	Greeter *Greeter
}

func NewMessage(g *Greeter) *Message {
	return &Message{Greeter: g}
}

func (m *Message) Print() {
	log.Println(m.Greeter.Greet("Wire"))
}

func main() {
	msg := InitializeMessage()
	msg.Print()
}
