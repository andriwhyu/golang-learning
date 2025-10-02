package main

import (
	"os"
)

func main() {
	dbAddress := os.Getenv("DB_ADDR")
	dbConnection := NewDB(dbAddress)

	connection, err := dbConnection.Connect()
	if err != nil {
		panic(err)
	}

	_ = connection
}
