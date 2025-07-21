package db

import (
	"context"
	"database/sql"
	"time"
)

func New(addr string, maxIdleConns, maxOpenConns int, maxIdleTime time.Duration) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)

	if err != nil {
		return nil, err
	}

	// set connection properties
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(maxIdleTime)

	ctx, timeout := context.WithTimeout(context.Background(), 5*time.Second)
	defer timeout()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
