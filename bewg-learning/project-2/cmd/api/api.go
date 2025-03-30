package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)

	return mux
}

func (app *application) run(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Printf("Starting server at %s\n", srv.Addr)

	return srv.ListenAndServe()
}
