package main

import (
	"fmt"
	"io"
	"slices-practice/exercise-1/24-fix-the-memory-leak/api"
)

func main() {
	// reports the initial memory usage
	api.Report()

	// returns a slice with 10 million elements.
	// it allocates 65 MB of memory space.
	millions := api.Read()

	// -----------------------------------------------------
	// ✪ ONLY CHANGE THE CODE IN THIS AREA ✪
	//last10 := make([]int, 10)
	//copy(last10, millions[len(millions)-10:])

	last10 := millions[len(millions)-10:]

	fmt.Printf("\nLast 10 elements: %d\n\n", last10)

	// ✪ ONLY CHANGE THE CODE IN THIS AREA ✪
	// -----------------------------------------------------
	millions = millions[0:1:1]
	millions = append(millions, 2)

	api.Report()

	// don't worry about this code.
	_, err := fmt.Fprintln(io.Discard, millions[0])
	if err != nil {
		return
	}

}

/*
	Learning point:
	- To garbage collection a slice, you need to make sure the backing array has been replaced. Example:
		abc := rand.Perm(2 << 2) // It will return a slice with 10 million int number (65 MB).
		abc = abc[0:1:1] // even though we've changed the element into 1, it still will calculate as 65MB since it still use the same backing array.
		abc = append(abc, 2) // it will create new backing array. The garbage collection will clear the previous backing array.
	- Processing a huge data with slice should be wise, otherwise even if we only use the 10 elements of the slice it will consume memory depends on the backing array instead of the slice itself.
*/
