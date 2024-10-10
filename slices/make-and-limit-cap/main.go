package main

import (
	"github.com/inancgumus/prettyslice"
	"strings"
)

func main() {
	/*
		learning Notes:
		- The full parameter of slice is [START:STOP:CAP]
		- CAP parameter used to limit the slice growth. One of the use case we want to limit the cap is to prevent the change/append of slice not impact the current backing array.
		- make have parameter make(type, len, [opt]cap); make([]string, 0, 5)
	*/
	prettyslice.MaxPerLine = 10
	prettyslice.PrintBacking = true

	toyStoryChar := []string{"woody", "buzz", "rex"}

	// 1st case without using make.
	//var upperChar []string
	//
	//prettyslice.Show("Init upperChar", upperChar)
	//
	//for _, val := range toyStoryChar {
	//	upperChar = append(upperChar, strings.ToUpper(val))
	//	prettyslice.Show("appended upperChar", upperChar)
	//}

	/*
		Summary:
		- The upperChar slice will create 3 different backing array.
		- Remember, every new backing array it means it allocate new memory and extra resources to copy the elements to the new memory.
	*/

	// 2nd case. Using make to allocate proper slices
	upperChar := make([]string, 0, cap(toyStoryChar))
	prettyslice.Show("Init upperChar", upperChar)

	for _, val := range toyStoryChar {
		upperChar = append(upperChar, strings.ToUpper(val))
		prettyslice.Show("appended upperChar", upperChar)
	}

	/*
		Summary:
		- The upperChar slice only create 1 backing array.
		- len=0 means if we append the slice, it will start from the first element.
	*/

}
