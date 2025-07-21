package main

import "database/sql"

type User struct {
	ID   int
	Name string
}

type Store interface {
	GetByID(int) (*User, error)
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetByID(id int) (*User, error) {
	row := r.db.QueryRow(`SELECT id, name FROM users WHERE id = $1`, id)

	var user User
	err := row.Scan(&user.ID, &user.Name)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
