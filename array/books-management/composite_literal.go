package main

import "fmt"

// Practice composite literal array

const (
	winter    = 1
	summer    = 3
	totalBook = winter + summer
)

func main() {
	libraryBooks := [totalBook]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	fmt.Printf("%#v\n", libraryBooks)
}
