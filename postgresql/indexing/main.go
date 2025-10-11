package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	dbAddress := os.Getenv("DB_ADDR")
	generatedDataCount, err := strconv.Atoi(os.Getenv("GENERATED_DATA_COUNT"))

	// set default number of generated data
	if err != nil {
		generatedDataCount = 5
	}

	// establish connection
	dbConnection := NewDB(dbAddress)
	connection, err := dbConnection.Connect()
	if err != nil {
		panic(err)
	}

	// running seeding
	seed := NewSeed(connection, generatedDataCount)
	err = seed.RunSeeding()
	if err != nil {
		fmt.Println("seeding automation is failed")
		panic(err)
	}
}
