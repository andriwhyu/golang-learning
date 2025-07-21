package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	dataDummy := map[string]string{
		"status":  "ok",
		"version": version,
		"env":     app.config.env,
	}

	err := app.jsonResponse(w, http.StatusOK, dataDummy)
	if err != nil {
		app.internalServerErrorLogger(w, r, err)
	}
}
