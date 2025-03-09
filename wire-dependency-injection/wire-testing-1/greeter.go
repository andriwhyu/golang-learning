package main

import "fmt"

type Greeter struct {
	Config *Config
}

func NewGreater(c *Config) *Greeter {
	return &Greeter{Config: c}
}

func (g *Greeter) Greet(name string) string {
	return fmt.Sprintf("%s, %s!", g.Config.Greeting, name)
}
