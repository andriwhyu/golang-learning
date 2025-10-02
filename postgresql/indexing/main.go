package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	dbAddress := os.Getenv("DB_ADDR")
	generatedDataCount, err := strconv.Atoi(os.Getenv("GENERATED_DATA_COUNT"))
	if err != nil {
		generatedDataCount = 5
	}

	dbConnection := NewDB(dbAddress)

	connection, err := dbConnection.Connect()
	if err != nil {
		panic(err)
	}

	seed := NewSeed(connection, generatedDataCount)
	err = seed.RunSeeding()
	if err != nil {
		fmt.Println("seeding automation is failed")
		panic(err)
	}
}
