package main

import (
	s "github.com/inancgumus/prettyslice"
	"sort"
)

func main() {
	/*
		Learning note:
		- Backing array is an array that created when slice literal was defined. This is done in the background.
		- When slicing an array/slice literal, golang not create a new array in memory, but it referred to backing array.
		  It means, when the array change, the slice value also change.
		- Each slice literal defined, it will create 1 backing array.
		- To separate a slice with the backing array, we need to declare a new slice literal, or append and empty slices with the new value.
	*/

	// Define a literal array
	//ages := []int{
	//	15, 31, 20,
	//}
	//fmt.Printf("first definition: %#v\n", ages)

	// change the 'backing' array value and print the origin ages
	//agesCopy := ages
	//agesCopy[0] = 10
	//fmt.Printf("backing array changes: %#v\n", ages) // it impacted to the change

	// adjust 'backing' array to create new array in memory
	//agesCopyNew := append([]int(nil), ages...)
	//fmt.Println()
	//fmt.Printf("agesCopyNew: %#v\n", agesCopyNew)
	//agesCopyNew[1] = 35
	//fmt.Printf("agesCopyNew adjusted: %#v\n", agesCopyNew)
	//fmt.Printf("ages original: %#v\n", ages)

	// short the element of grades
	grades := [...]float64{40, 10, 20, 50, 60, 70}
	front := grades[:3]

	sort.Float64s(front)

	s.PrintBacking = true
	s.MaxPerLine = 7
	s.Show("grades", grades[:])
	s.Show("front", front)
	s.Show("get middle elements", grades[2:4])

}
