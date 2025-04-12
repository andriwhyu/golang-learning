package main

import (
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/internal/env"
	"log"
)

func main() {
	cfg := &config{
		addr: env.GetStringVar("ADDR", ":8080"),
	}

	app := &application{
		config: *cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
