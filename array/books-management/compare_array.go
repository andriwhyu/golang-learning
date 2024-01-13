package main

import "fmt"

/*
Rule of thumb:
To compare arrays make sure it have the same type. Ex:
A := [5]string{} == B := [5]string{} (✅)
A := [5]string{} == B := [5]int{} (⛔️)
A := [5]string{} == B := [6]string{} (⛔️)
A := [5]int{} == B := [5]int64{} (⛔️)
*/
func main() {
	// Show number of books on each shelf
	shelfA := [3]int{3, 6, 9}
	shelfB := [3]int{3, 6, 9}

	fmt.Printf("Shelf A: %#v\n", shelfA)
	fmt.Printf("Shelf B: %v\n", shelfB)
	fmt.Println("Comparison result:", shelfA == shelfB)

	shelfA = [3]int{3, 6}
	shelfB = [3]int{3, 6, 9}

	fmt.Printf("Shelf A: %v\n", shelfA)
	fmt.Printf("Shelf B: %v\n", shelfB)
	fmt.Println("Comparison result:", shelfA == shelfB)

	shelfA = [3]int{3, 6}
	shelfB = [3]int{3, 6, 9}

	fmt.Printf("Shelf A: %v\n", shelfA)
	fmt.Printf("Shelf B: %v\n", shelfB)
	fmt.Println("Comparison result:", shelfA == shelfB)

}
