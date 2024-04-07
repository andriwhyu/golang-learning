package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		Objective:
		- Generate a random number and store only the unique number.
		- The size of the data structure depends on user input

		Learning notes:
		- An array can be converted into slice by using `:`, e.g abc[:]
	*/

	//	Create for to generate the random number
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	occurance := 10
	var allNumber []int

generate_random:
	for i := 0; i < occurance; {
		randNumber := randomizer.Intn(occurance + 1)
		fmt.Printf("%d ", randNumber)

		// Check if the generated number have been created before, if yes then reproduce the random number.
		for _, val := range allNumber {
			if randNumber == val {
				continue generate_random
			}

		}
		allNumber = append(allNumber, randNumber)
		i++
	}

	fmt.Println()
	fmt.Printf("Generated number: %#v\n", allNumber)

}
