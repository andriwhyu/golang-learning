package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	usage = `

Usage:
go run main.go [string]
go run main.go sTaY
`
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Print("Tell me a book title.", usage)
		return
	}

	isEverMatch := false
	bookTitle := args[0]
	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	fmt.Println("Search Results:")

	for _, val := range books {
		isMatch := strings.Contains(strings.ToLower(val), strings.ToLower(bookTitle))

		if isMatch {
			fmt.Println("+", val)
			isEverMatch = true
		}
	}

	if !isEverMatch {
		fmt.Printf("We don't have the book: \"%s\"\n", bookTitle)
	}

}
