package main

import (
	"fmt"
)

func main() {
	/*
		Learning note:
		- Append to a slice that only in specific length will impact the original slices / array
		- The differences between slices and indexing is the return value. Slices will always return slices, meanwhile indexing will return
		  the original value, e.g. int, string.

	*/
	slicesAlphabeth := []string{"h", "e", "l", "l", "o"}
	slicePart := append(slicesAlphabeth[1:3], "!")

	fmt.Printf("%v\n", slicesAlphabeth)
	fmt.Printf("%v\n", slicePart)

	realMadridPlayer := []string{"Camavinga", "Tchouameni", "Bellingham", "Valverde"}

	fmt.Printf("slicing: %T %[1]q\n", realMadridPlayer[1:2])
	fmt.Printf("indexing: %T %[1]q\n", realMadridPlayer[1])

}
