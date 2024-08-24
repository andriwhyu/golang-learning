package main

import (
	"fmt"
	s "github.com/inancgumus/prettyslice"
	"unsafe"
)

type collection []string

func change(data collection) {
	// pass a value to a function is make a copy of the variable locally (in func)
	data[2] = "marvelous"

	s.Show("change func data:", data)
	fmt.Printf("change func data: %p\n", data)
}

func main() {
	/*
		Learning note:
		- Every slice will behind the scene have a tiny data structure called slice header.
		- Slice header have 3 components; pointer, length, capacity.
		- Pointer is the memory address of the first element of the backing array.
		- Length is the length of the slice.
		- Capacity is calculated from the first element of the slice to the last element of the backing array.
		- Size of slice = size of 'slice header'. Size of slice header will always the same for each slice.
		- A string occupied 16 bytes in machine 64bit.
		- The capacity of slice depends on the slice pointer.
		- When append a slice with full capacity condition, go with allocate a new backing array with a larger capacity
		  more than it required. It kinda forecast the future based on additional elements. Because reallocating is a resource
		  consuming.
		- Append works by appending the last element of slice by looking the length of the slice.
	*/
	s.PrintElementAddr = true

	data := collection{"slices", "are", "awesome", "period"}
	change(data)

	s.Show("main data:", data)
	fmt.Printf("main data: %p\n", data)

	// size of array vs slice
	array := [...]string{"slices", "are", "awesome", "period"}

	fmt.Printf("array's size: %d bytes\n", unsafe.Sizeof(array))
	fmt.Printf("slices's size: %d bytes\n", unsafe.Sizeof(data))

	//var testArr [1000]int64
	//
	//testArr2 := testArr
	//sliceArr := testArr2[:]
	//_ = sliceArr
	//fmt.Printf("testArr size: %d bytes\n", unsafe.Sizeof(testArr))

	// capacity learning
	//words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
	//words = words[:0]
	s.PrintBacking = true
	s.MaxPerLine = 10
	//s.Show("words", words)
	//s.Show("words1", words[:cap(words)])

	s.Show("data before append", data)
	data = append(data, "wow", "yes baby", "go")
	s.Show("data after append", data)

}
