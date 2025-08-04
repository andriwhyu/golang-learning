package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int) (*Post, error)
		DeleteByID(context.Context, int) error
		UpdateByID(context.Context, int, *Post) error
	}

	Users interface {
		Create(context.Context, *User) error
		GetByID(context.Context, int) (*User, error)
	}

	Comments interface {
		Create(context.Context, *Comment) error
		GetByPostID(context.Context, int) (*[]Comment, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{
			db: db,
		},
		Users: &UserStore{
			db: db,
		},
		Comments: &CommentStore{
			db: db,
		},
	}
}
