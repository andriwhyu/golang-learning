package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/constants"
	"github.com/lib/pq"
)

type Post struct {
	ID           int        `json:"id"`
	Title        string     `json:"title"`
	Content      string     `json:"content"`
	Tags         []string   `json:"tags"`
	UserID       int        `json:"user_id"`
	CreatedAt    string     `json:"created_at"`
	UpdatedAt    string     `json:"updated_at"`
	Version      int        `json:"version"`
	Comments     *[]Comment `json:"comments"`
	CommentCount int        `json:"comment_count"`
}

type PostFeed struct {
	Post Post `json:"post"`
	User User `json:"user"`
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

func (ps *PostStore) GetPostFeed(ctx context.Context, id int, pagination Pagination) ([]*PostFeed, error) {
	// the query will get all posts for the selected user or all posts from all user that has been followed by the selected user
	query := fmt.Sprintf(`
		SELECT
			p.id, p.title, p.content, p.tags, p.user_id, p.created_at, p.updated_at, p.version,
			u.id, u.username, u.first_name, u.last_name, u.email, u.created_at,
			COUNT(p.id) AS total_comment
		FROM posts p
		LEFT JOIN comments c ON p.id = c.post_id
		JOIN users u ON u.id = p.user_id
		WHERE p.user_id = $1 OR p.user_id IN (
			SELECT user_id FROM followers WHERE follower_id = $1
		)
		GROUP BY p.id, u.id
		ORDER BY p.created_at %s
		LIMIT $2 OFFSET $3;
	`, pagination.Sort)

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeout)
	defer cancel()

	rows, err := ps.db.QueryContext(ctx, query, id, pagination.Limit, pagination.Offset)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		errDefer := rows.Close()
		if errDefer != nil && err == nil {
			err = errDefer
		}
	}(rows)

	var postFeeds []*PostFeed

	for rows.Next() {
		var (
			post Post
			user User
		)

		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			pq.Array(&post.Tags),
			&post.UserID,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.Version,
			&user.ID,
			&user.Username,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.CreatedAt,
			&post.CommentCount,
		)
		postFeed := PostFeed{Post: post, User: user}
		postFeeds = append(postFeeds, &postFeed)

		if err != nil {
			return nil, err
		}
	}

	return postFeeds, nil
}
