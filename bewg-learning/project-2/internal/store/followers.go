package store

import (
	"context"
	"database/sql"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/constants"
)

type Followers struct {
	UserID     int    `json:"user_id" validate:"required"`
	FollowerID int    `json:"follower_id" validate:"required"`
	CreatedAt  string `json:"created_at"`
}

type FollowerStore struct {
	db *sql.DB
}

func (fs *FollowerStore) Follow(ctx context.Context, userID int, followerID int) error {
	query := `
		INSERT INTO followers (user_id, follower_id)
		VALUES ($1, $2) ON CONFLICT DO NOTHING
	`

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeout)
	defer cancel()

	_, err := fs.db.ExecContext(ctx, query, userID, followerID)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FollowerStore) Unfollow(ctx context.Context, userID, followerID int) error {
	query := `
		DELETE FROM followers WHERE user_id = $1 AND follower_id = $2
	`
	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeout)
	defer cancel()

	_, err := fs.db.ExecContext(ctx, query, userID, followerID)
	if err != nil {
		return err
	}

	return nil
}
