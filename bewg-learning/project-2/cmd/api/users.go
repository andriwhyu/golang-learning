package main

import (
	"errors"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/constants"
	"net/http"
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := getParamID(r, "userID")
	if err != nil {
		app.badRequestErrorLogger(w, r, err)
		return
	}

	user, err := app.store.Users.GetByID(ctx, userID)
	if err != nil {
		switch {
		case errors.Is(err, constants.ErrDataNotFoundByID):
			app.statusNotFoundErrorLogger(w, r, err)
		default:
			app.internalServerErrorLogger(w, r, err)
		}
		return
	}

	err = app.jsonResponse(w, http.StatusOK, user)
	if err != nil {
		app.internalServerErrorLogger(w, r, err)
	}
}
