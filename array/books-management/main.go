package main

import "fmt"

const (
	winterBooksCount = 1
	summerBooksCount = 3
	booksCount       = winterBooksCount + summerBooksCount
)

func main() {
	var books [booksCount]string
	var winterBooks [winterBooksCount]string
	var summerBooks [summerBooksCount]string
	var publishedBooks [len(books)]bool

	books[0] = "Kafka's Revenge"
	books[1] = "Stay Golden"
	books[2] = "Everythingship"
	books[3] += books[0] + " 2nd Edition"

	// For range actually do copied the object. Thus, if
	// we update the element it won't effect.

	winterBooks[0] = books[0]
	for i, _ := range summerBooks {
		summerBooks[i] = books[i+1]
	}

	publishedBooks[0] = true
	publishedBooks[len(books)-1] = true

	fmt.Printf("All books: %#v\n", books)
	fmt.Printf("Winter books: %#v\n", winterBooks)
	fmt.Printf("Summer books: %#v\n", summerBooks)

	fmt.Println("Published Books:")
	for i, v := range publishedBooks {
		if v {
			fmt.Printf("+ %s\n", books[i])
		}
	}
}
