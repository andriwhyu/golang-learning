package main

import (
	"net/http"
)

func (app *application) getFeedHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	postFeed, err := app.store.Posts.GetPostFeed(ctx, 3)
	if err != nil {
		app.internalServerErrorLogger(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, postFeed); err != nil {
		app.internalServerErrorLogger(w, r, err)
	}
}
