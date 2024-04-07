package main

import "fmt"

func main() {
	/*
		Learning note:
		- Append will be used to add element in a slice.
		- Append func will return a new slice with attached the new data (appended value).
		- Add `<slice>...` to append a slice in another slice.
	*/

	book := []string{"Avengers reborn", "Avatar new adventure"}

	// Append new book
	book = append(book, "Captain America")
	fmt.Printf("1. Regular append: %#v\n", book)

	//	Append a slice wit another slice
	hotDealsBook := []string{"Remember me", "Cocktail vibes"}
	book = append(book, hotDealsBook...)
	fmt.Printf("2. Slice append: %#v\n", book)

	//	Multiple append to slice
	executiveBook := append(hotDealsBook, "Gold Investment", "Coal energy")
	fmt.Printf("3. Executive book: %#v\n", executiveBook)
}
