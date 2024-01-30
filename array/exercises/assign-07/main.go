package main

import (
	"fmt"
	"strings"
)

func main() {
	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
	}

	upper, lower := books, books

	for i, val := range upper {
		upper[i] = strings.ToUpper(val)
		lower[i] = strings.ToLower(val)
	}

	fmt.Printf("%q\n", books)
	fmt.Printf("%q\n", upper)
	fmt.Printf("%q\n", lower)

}
