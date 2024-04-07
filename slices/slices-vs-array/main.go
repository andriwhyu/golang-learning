package main

import "fmt"

func main() {
	/*
		Learning note:
		- Uninitialize and empty slice is different. If want to check whether the slice empty or not, check the len.
		- Empty/uninitialize slice cannot be assign in any index, e.g bookSliceInit[0] = "abc". It will return runtime error
		- Comparing two different slices only can be done by compare each element. Unlike array, we can't compare two slices directly.

	*/
	// Variable declaration
	var bookSlice []string  // Required field: data type.
	var bookArray [2]string // required filed: data type, array size.

	bookSliceInit := []string{}

	fmt.Printf("book slice: %#v\n", bookSlice)     // uninitialize result
	fmt.Printf("book slice: %#v\n", bookSliceInit) // empty slice
	fmt.Printf("book array: %#v\n", bookArray)     // Empty array

	bookSummer := []string{"avenger I", "avenger reborn"}
	bookWinter := []string{"Iron man new life"}

	fmt.Printf("Release books in summer: %#v\n", bookSummer)
	fmt.Printf("Release books in winter: %#v\n", bookWinter)

	// Example of invalid comparison
	//if bookSummer == bookWinter {
	//	fmt.Println("Hi there")
	//}

	//	Simple example of the working comparison.
	//summer:
	//	for _, summer := range bookSummer {
	//		for _, winter := range bookWinter {
	//			if summer != winter {
	//				break summer
	//			}
	//		}
	//	}

}
