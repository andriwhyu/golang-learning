package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	addr string
}

var user []User

func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload User

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validateInput(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user = append(user, payload)

	w.WriteHeader(http.StatusCreated)
}

func validateInput(u *User) error {
	if u.FirstName == "" || u.LastName == "" {
		return errors.New("there is missing field")
	}

	for _, eachUser := range user {
		if eachUser.FirstName == u.FirstName && eachUser.LastName == u.LastName {
			return errors.New("user has been exist")
		}
	}
	return nil
}
