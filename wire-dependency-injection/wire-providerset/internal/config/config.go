package config

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewConfig)

type Config struct {
	Greating string
}

func NewConfig() *Config {
	return &Config{Greating: "Hi"}
}
