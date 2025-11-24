package main

import (
	"database/sql"
	"fmt"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/internal/db"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/internal/env"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/internal/store"
	"os"
	"strconv"
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

	generatedDataCount, err := strconv.Atoi(os.Getenv("GENERATED_DATA_COUNT"))

	// set default number of generated data
	if err != nil {
		generatedDataCount = 100
	}

	fmt.Printf("Generated data count: %d\n", generatedDataCount)

	storageObj := store.NewStorage(dbConnection)
	db.Seed(storageObj, generatedDataCount)
}
