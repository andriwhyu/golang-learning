package main

import (
	"database/sql"
	"log"

	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/internal/db"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/internal/env"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/internal/store"
)

const (
	version = "0.0.1"
)

func main() {
	cfg := &config{
		addr: env.GetStringVar("ADDR", ":8080"),
		env:  env.GetStringVar("ENV", "development"),
		dbConfig: dbConfig{
			addr:         env.GetStringVar("DB_ADDR", "postgresql://admin:adminpassword@localhost:5432/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetIntVar("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetIntVar("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetDurationStringVar("DB_MAX_IDLE_TIME", "15m"), // valid pattern e.g. 2h15m10s, 2h, 15m, 10s
		},
	}

	dbConnection, err := db.New(cfg.dbConfig.addr, cfg.dbConfig.maxOpenConns, cfg.dbConfig.maxIdleConns, cfg.dbConfig.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer func(dbConnection *sql.DB) {
		err := dbConnection.Close()
		if err != nil {
			log.Panic(err)
		}
	}(dbConnection)

	log.Println("Database connection established")

	appStorage := store.NewStorage(dbConnection)

	app := &application{
		config: *cfg,
		store:  appStorage,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
