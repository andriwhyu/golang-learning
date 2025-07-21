package store

import (
	"context"
	"database/sql"
	"errors"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/constants"
)

type Comment struct {
	ID        int    `json:"id"`
	Content   string `json:"content" validate:"required"`
	UserID    int    `json:"user_id" validate:"required"`
	PostID    int    `json:"post_id" validate:"required"`
	CreatedAt string `json:"created_at"`
	User      *User  `json:"user"`
}

type CommentStore struct {
	db *sql.DB
}

func (cm *CommentStore) Create(ctx context.Context, comment *Comment) error {
	query := `
		INSERT INTO comments(content, users_id, posts_id)
		VALUES ($1, $2, $3) RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeout)
	defer cancel()

	err := cm.db.QueryRowContext(
		ctx,
		query,
		comment.Content,
		comment.UserID,
		comment.PostID,
	).Scan(&comment.ID, &comment.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}
func (cm *CommentStore) GetByPostID(ctx context.Context, postID int) (*[]Comment, error) {
	query := `
		SELECT 
		    c.id, c.user_id, c.post_id, c.content, c.created_at,
		    u.id, u.username, u.email, u.first_name, u.last_name, u.created_at
		FROM comments c
		INNER JOIN users u ON u.id = c.user_id
		WHERE c.post_id = $1
		ORDER BY c.created_at DESC
	`

	var comments []Comment

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeout)
	defer cancel()

	rows, err := cm.db.QueryContext(ctx, query, postID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, constants.ErrDataNotFoundByID
		default:
			return nil, err
		}
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var (
			comment Comment
			user    User
		)

		err = rows.Scan(
			&comment.ID, &comment.UserID, &comment.PostID, &comment.Content, &comment.CreatedAt,
			&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt,
		)

		comment.User = &user

		if err != nil {
			return &comments, err
		}

		comments = append(comments, comment)
	}

	return &comments, nil
}
