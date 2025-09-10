package main

import (
	"context"
	"errors"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/constants"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/internal/store"
	"net/http"
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)

	err := app.jsonResponse(w, http.StatusOK, user)
	if err != nil {
		app.internalServerErrorLogger(w, r, err)
	}
}

func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser := getUserFromContext(r)
	userID := 1 // will replace later once the auth done

	ctx := r.Context()
	err := app.store.Followers.Follow(ctx, userID, followerUser.ID)

	if err != nil {
		app.internalServerErrorLogger(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerErrorLogger(w, r, err)
		return
	}
}

func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser := getUserFromContext(r)
	userID := 1 // will replace later once the auth done

	ctx := r.Context()
	err := app.store.Followers.Unfollow(ctx, userID, followerUser.ID)

	if err != nil {
		app.internalServerErrorLogger(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerErrorLogger(w, r, err)
		return
	}
}

func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := getParamID(r, "userID")
		if err != nil {
			app.badRequestErrorLogger(w, r, err)
			return
		}

		ctx := r.Context()
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

		ctx = context.WithValue(ctx, constants.UserCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromContext(r *http.Request) *store.User {
	return r.Context().Value(constants.UserCtx).(*store.User)
}
