package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DBConnection struct {
	address string
}

func NewDB(address string) *DBConnection {
	return &DBConnection{address: address}
}

func (db *DBConnection) Connect() (*sql.DB, error) {
	dbConn, err := sql.Open("postgres", db.address)
	if err != nil {
		fmt.Println("Error connecting to database")
		return nil, err
	}

	if err := dbConn.Ping(); err != nil {
		fmt.Println("Cannot ping to database")
		return nil, err
	}

	return dbConn, nil
}
