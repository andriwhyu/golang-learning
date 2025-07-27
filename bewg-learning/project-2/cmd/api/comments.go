package main

import (
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/constants"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/internal/store"
	"net/http"
)

type createCommentRequest struct {
	Content string `json:"content" validate:"required,max=1000"`
}

func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	var reqData createCommentRequest

	if err := readJson(w, r, &reqData); err != nil {
		app.badRequestErrorLogger(w, r, err)
		return
	}

	if err := constants.Validator.Struct(reqData); err != nil {
		app.badRequestErrorLogger(w, r, err)
		return
	}

	ctx := r.Context()
	post := getPostFromContext(r)

	comment := &store.Comment{
		PostID:  post.ID,
		Content: reqData.Content,
		// TODO: change this with the user ID of commented user.
		UserID: 1,
	}

	if err := app.store.Comments.Create(ctx, comment); err != nil {
		app.internalServerErrorLogger(w, r, err)
	}
}
