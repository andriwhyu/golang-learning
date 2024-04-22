package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Get numbers from command-line
	userInput := os.Args[1:]

	if len(userInput) == 0 {
		fmt.Println("provide a few numbers")
		return
	}

	// Convert the input into a number
	var output []int

	for _, e := range userInput {
		intResult, err := strconv.Atoi(e)

		if err != nil {
			continue
		}

		output = append(output, intResult)
	}

	sort.Ints(output)
	fmt.Printf("%v\n", output)
}
