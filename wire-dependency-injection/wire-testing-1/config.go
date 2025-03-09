package main

type Config struct {
	Greeting string
}

func NewConfig() *Config {
	return &Config{Greeting: "Hello"}
}
