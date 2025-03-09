//wire.go
//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// InitializeMessage is the Wire injector function.
func InitializeMessage() *Message {
	wire.Build(NewMessage, NewConfig, NewGreater)
	return nil
}
