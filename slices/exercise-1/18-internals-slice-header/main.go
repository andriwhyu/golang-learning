package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"
)

const size = 1e7

func main() {
	// don't worry about this code.
	// it stops the garbage collector: prevents cleaning up the memory.
	// see the link if you're curious:
	// https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)
	debug.SetGCPercent(-1)

	// run the program to see the initial memory usage.
	report("initial memory usage")

	// 1. allocate an array with 10 million int elements
	//    the array's size will be equal to ~80MB
	//    hint: use the `size` constant above.
	var newArr [size]int
	_ = newArr
	// 2. print the memory usage (use the report func).
	report("after declaring an array")

	// 3. copy the array to a new array.
	cpArray := newArr

	// 4. print the memory usage
	report("after copying the array")

	// 5. pass the array to the passArray function
	passArray(cpArray)

	// 6. convert one of the arrays to a slice
	newSlice := cpArray[:]
	_ = newSlice

	// 7. slice only the first 1000 elements of the array
	sliceOnly1000 := cpArray[:1000]
	_ = sliceOnly1000

	// 8. slice only the elements of the array between 1000 and 10000
	sliceAgain := cpArray[1000:10000]
	_ = sliceAgain

	// 9. print the memory usage (report func)
	report("after slicing")

	// 10. pass the one of the slices to the passSlice function
	passSlice(newSlice)

	// 11. print the sizes of the arrays and slices
	//     hint: use the unsafe.Sizeof function
	//     see more here: https://golang.org/pkg/unsafe/#Sizeof
	fmt.Printf("\nArray's size : %d bytes\n", unsafe.Sizeof(newArr))
	fmt.Printf("Array2's size: %d bytes\n", unsafe.Sizeof(cpArray))
	fmt.Printf("Slice1's size: %d bytes\n", unsafe.Sizeof(newSlice))
	fmt.Printf("Slice2's size: %d bytes\n", unsafe.Sizeof(sliceOnly1000))
	fmt.Printf("Slice3's size: %d bytes\n", unsafe.Sizeof(sliceAgain))
}

// passes [size]int array — about 80MB!
//
// observe that passing an array to a function (or assigning it to a variable)
// affects the memory usage dramatically
func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// only passes 24-bytes of slice header
//
// observe that passing a slice doesn't affect the memory usage
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

// reports the current memory usage
// don't worry about this code
func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
