package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {

	_, err := fmt.Fprintf(w, "hello\n")
	if err != nil {
		return
	}
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			_, err := fmt.Fprintf(w, "%v: %v\n", name, h)
			if err != nil {
				return
			}
		}
	}
}

func main() {

	//http.HandleFunc("/hello", hello)
	//http.HandleFunc("/headers", headers)
	//
	//err := http.ListenAndServe(":8090", nil)

	//create custom mux
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/headers", headers)

	err := http.ListenAndServe(":8090", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
