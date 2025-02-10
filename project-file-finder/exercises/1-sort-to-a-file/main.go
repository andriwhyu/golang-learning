package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	sort.Strings(args)

	// calculate number of length
	var lenArgs int
	for _, val := range args {
		lenByteVal := len([]rune(val)) + 1
		lenArgs += lenByteVal
	}

	argsByte := make([]byte, 0, lenArgs)
	for _, val := range args {
		argsByte = append(argsByte, val...)
		argsByte = append(argsByte, "\n"...)
	}

	fmt.Printf("total bytes: %d\n", cap(argsByte))

	err := os.WriteFile("out.txt", argsByte, 0644)
	if err != nil {
		fmt.Printf("Error on writing file: %s\n", err)
		return
	}
}
