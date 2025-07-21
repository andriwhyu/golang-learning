package main

import (
	"encoding/json"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/constants"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func writeJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func readJson(w http.ResponseWriter, r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)

	r.Body = http.MaxBytesReader(w, r.Body, constants.MaxBytesRequest) // limit the request body
	decoder.DisallowUnknownFields()                                    // disallow unknown fields

	return decoder.Decode(data)
}

func writeErrorJson(w http.ResponseWriter, status int, msg string) error {
	type envelope struct {
		Error string `json:" "`
	}

	return writeJson(w, status, &envelope{Error: msg})
}

func (app *application) jsonResponse(w http.ResponseWriter, status int, data any) error {
	type envelope struct {
		Data any `json:"data"`
	}
	return writeJson(w, status, &envelope{Data: data})
}

func getParamID(r *http.Request, param string) (int, error) {
	idStr := chi.URLParam(r, param)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	return id, nil
}
