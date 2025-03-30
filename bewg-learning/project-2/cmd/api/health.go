package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("OK\n"))

	if err != nil {
		fmt.Println(err)
	}
}
