package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerErrorLogger(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("error found with protocol: %s path: %s error: %s\n", r.Method, r.URL, err.Error())

	err = writeErrorJson(w, http.StatusInternalServerError, "server encountered an error")
	if err != nil {
		return
	}
}

func (app *application) statusNotFoundErrorLogger(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("error found with protocol: %s path: %s error: %s\n", r.Method, r.URL, err.Error())

	err = writeErrorJson(w, http.StatusNotFound, "data not found")
	if err != nil {
		return
	}
}

func (app *application) conflictErrorLogger(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("error found with protocol: %s path: %s error: %s\n", r.Method, r.URL, err.Error())

	err = writeErrorJson(w, http.StatusConflict, "conflict operation")
	if err != nil {
		return
	}
}

func (app *application) badRequestErrorLogger(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("error found with protocol: %s path: %s error: %s\n", r.Method, r.URL, err.Error())

	err = writeErrorJson(w, http.StatusBadRequest, "invalid request")
	if err != nil {
		return
	}
}
