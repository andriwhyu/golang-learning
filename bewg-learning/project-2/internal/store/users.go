package store

import (
	"database/sql"
	"errors"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/constants"
	"golang.org/x/net/context"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (us *UserStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users(username, email, first_name, last_name, password) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeout)
	defer cancel()

	err := us.db.QueryRowContext(
		ctx, query, user.Username, user.Email,
		user.FirstName, user.LastName, user.Password,
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (us *UserStore) GetByID(ctx context.Context, id int) (*User, error) {
	query := `
		SELECT 
			id, username, email, first_name, last_name, created_at
		FROM users
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeout)
	defer cancel()

	var user User
	err := us.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, constants.ErrDataNotFoundByID
		default:
			return nil, err
		}
	}

	return &user, nil
}
