package main

import (
	"net/http"

	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/constants"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/internal/store"
)

func (app *application) getFeedHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pagination := &store.Pagination{}

	pagination, err := pagination.Parse(r)
	if err != nil {
		app.badRequestErrorLogger(w, r, err)
		return
	}

	err = constants.Validator.Struct(pagination)
	if err != nil {
		app.badRequestErrorLogger(w, r, err)
		return
	}

	postFeed, err := app.store.Posts.GetPostFeed(ctx, 3, *pagination)
	if err != nil {
		app.internalServerErrorLogger(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, postFeed); err != nil {
		app.internalServerErrorLogger(w, r, err)
	}
}
