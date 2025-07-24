package main

import (
	"database/sql"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/internal/db"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/internal/env"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/internal/store"
)

func main() {
	dbAddr := env.GetStringVar("DB_ADDR", "postgresql://admin:adminpassword@localhost:5432/socialnetwork?sslmode=disable")
	dbMaxOpenConns := 3
	dbMaxIdleConns := 2
	dbMaxIdleTime := env.GetDurationStringVar("DB_MAX_IDLE_TIME", "15m")

	dbConnection, err := db.New(dbAddr, dbMaxOpenConns, dbMaxIdleConns, dbMaxIdleTime)
	if err != nil {
		panic(err)
	}

	defer func(dbConnection *sql.DB) {
		err := dbConnection.Close()
		if err != nil {
			panic(err)
		}
	}(dbConnection)

	storageObj := store.NewStorage(dbConnection)
	db.Seed(storageObj)
}
