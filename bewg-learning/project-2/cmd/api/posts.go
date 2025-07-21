package main

import (
	"context"
	"errors"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/constants"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/internal/store"
	"net/http"
)

type createPostRequest struct {
	Title   string   `json:"title" validate:"required,max=255"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}

type updatePostRequest struct {
	Title   *string   `json:"title" validate:"omitempty,max=255"`
	Content *string   `json:"content" validate:"omitempty,max=1000"`
	Tags    *[]string `json:"tags" validate:"omitempty"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var reqData createPostRequest

	if err := readJson(w, r, &reqData); err != nil {
		app.badRequestErrorLogger(w, r, err)
		return
	}

	if err := constants.Validator.Struct(reqData); err != nil {
		app.badRequestErrorLogger(w, r, err)
		return
	}

	post := &store.Post{
		Title:   reqData.Title,
		Content: reqData.Content,
		Tags:    reqData.Tags,
		// TODO: change this with dynamic userID
		UserID: 1,
	}
	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.internalServerErrorLogger(w, r, err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	post := getPostFromContext(r)

	comments, err := app.store.Comments.GetByPostID(ctx, post.ID)
	if err != nil {
		app.internalServerErrorLogger(w, r, err)
		return
	}

	post.Comments = comments

	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.internalServerErrorLogger(w, r, err)
	}
}

func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromContext(r)

	ctx := r.Context()
	err := app.store.Posts.DeleteByID(ctx, post.ID)
	if err != nil {
		switch {
		case errors.Is(err, constants.ErrDataNotFoundByID):
			app.statusNotFoundErrorLogger(w, r, err)
		default:
			app.internalServerErrorLogger(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	var payload updatePostRequest
	if err := readJson(w, r, &payload); err != nil {
		app.badRequestErrorLogger(w, r, err)
	}

	if err := constants.Validator.Struct(payload); err != nil {
		app.badRequestErrorLogger(w, r, err)
	}

	ctx := r.Context()
	post := getPostFromContext(r)

	if payload.Title != nil {
		post.Title = *payload.Title
	}

	if payload.Content != nil {
		post.Content = *payload.Content
	}

	if payload.Tags != nil {
		post.Tags = *payload.Tags
	}

	err := app.store.Posts.UpdateByID(ctx, post.ID, post)
	if err != nil {
		switch {
		case errors.Is(err, constants.ErrConflict):
			app.conflictErrorLogger(w, r, err)
		default:
			app.internalServerErrorLogger(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.internalServerErrorLogger(w, r, err)
	}
}

func (app *application) postContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		postID, err := getParamID(r, "postID")
		if err != nil {
			app.badRequestErrorLogger(w, r, err)
		}

		ctx := r.Context()
		post, err := app.store.Posts.GetByID(ctx, postID)
		if err != nil {
			switch {
			case errors.Is(err, constants.ErrDataNotFoundByID):
				app.statusNotFoundErrorLogger(w, r, err)
			default:
				app.internalServerErrorLogger(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, constants.PostCtx, post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getPostFromContext(r *http.Request) *store.Post {
	return r.Context().Value(constants.PostCtx).(*store.Post)
}
