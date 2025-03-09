package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := &http.Server{
		Addr: "8080",
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			time.Sleep(10 * time.Second)
			writer.Write([]byte("Hello, World!"))
		}),
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Server listening on: %s\n", srv.Addr)
	}()

}
