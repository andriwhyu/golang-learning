package store

import (
	"context"
	"database/sql"
	"errors"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/constants"
	"github.com/lib/pq"
)

type Post struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Tags      []string   `json:"tags"`
	UserID    int        `json:"user_id"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
	Version   int        `json:"version"`
	Comments  *[]Comment `json:"comments"`
}

type PostStore struct {
	db *sql.DB
}

func (ps *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts(title, content, tags, user_id) 
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeout)
	defer cancel()

	err := ps.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		pq.Array(post.Tags),
		post.UserID,
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (ps *PostStore) GetByID(ctx context.Context, id int) (*Post, error) {
	var post Post

	query := `
		SELECT 
		    id, title, content, tags, user_id, created_at, updated_at, version 
		FROM 
		    posts 
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeout)
	defer cancel()

	err := ps.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		pq.Array(&post.Tags),
		&post.UserID,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, constants.ErrDataNotFoundByID
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (ps *PostStore) DeleteByID(ctx context.Context, id int) error {
	query := `
		DELETE FROM posts
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeout)
	defer cancel()

	result, err := ps.db.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	res, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if res < 1 {
		return constants.ErrDataNotFoundByID
	}

	return nil
}

func (ps *PostStore) UpdateByID(ctx context.Context, id int, post *Post) error {
	query := `
		UPDATE posts
		SET title = $2, content = $3, tags = $4, version = $5 + 1
		WHERE id = $1 AND version = $5
		RETURNING version
	`

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeout)
	defer cancel()

	err := ps.db.QueryRowContext(
		ctx,
		query,
		id,
		post.Title,
		post.Content,
		pq.Array(post.Tags), post.Version,
	).Scan(&post.Version)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return constants.ErrConflict
		default:
			return err
		}
	}

	return nil
}
