package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"
const usage = "Please give me a word to search."

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]
	filter := [...]string{"and", "or", "was", "the", "since", "very"}

	if len(query) < 1 {
		fmt.Println(usage)
		return
	}
	// labels and other names do not share the same scope
	// this will work even though `queries` label exists
	//
	// var queries string
	// _ = queries

	// this label labels the parent loop below.
	// label's scope is the whole func main()
	//queries:
	//	for _, q := range query {
	//		q = strings.ToLower(q)
	//
	//	search:
	//		for i, w := range words {
	//			switch q {
	//			case "and", "or", "was", "the", "since", "very":
	//				break search
	//			}
	//
	//			if q == w {
	//				fmt.Printf("#%-2d: %q\n", i+1, w)
	//
	//				// find the first word then quit
	//				continue queries
	//			}
	//		}
	//	}
	//}
queries:
	for _, q := range query {
		q = strings.ToLower(q)

		// Search for filtering words
		for _, val := range filter {
			if q == val {
				break queries
			}
		}

		for i, word := range words {
			if q == word {
				fmt.Printf("#%-2d: %q\n", i+1, word)
				continue queries
			}
		}
	}

}
