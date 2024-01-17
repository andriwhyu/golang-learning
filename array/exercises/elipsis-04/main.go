package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		names     = [...]string{"Einstein", "Tesla", "Shepard"}
		distances = [...]int{50, 40, 75, 30, 125}
		data      = [...]uint{72, 69, 76, 76, 79}
		ratio     = [...]float64{3.14}
		alives    = [...]bool{false, true, true, false}
		zero      = [...]uint8{}
	)

	separator := "\n" + strings.Repeat("=", 20) + "\n"

	// Without loop range
	fmt.Print("\nnames", separator)
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %s\n", i, names[i])
	}
	fmt.Println()

	fmt.Println("distances\n====================")
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}
	fmt.Println()

	fmt.Println("data\n====================")
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}
	fmt.Println()

	fmt.Println("ratio\n====================")
	for i := 0; i < len(ratio); i++ {
		fmt.Printf("ratio[%d]: %.2f\n", i, ratio[i])
	}
	fmt.Println()

	fmt.Println("alives\n====================")
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}
	fmt.Println()

	fmt.Println("zero\n====================")
	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d]: %d\n", i, zero[i])
	}
	fmt.Println()

	//fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\nFOR RANGES\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Printf(`
%s
FOR RANGES
%[1]s

`, strings.Repeat("~", 30))

	// Range
	fmt.Println("names\n====================")
	for i, val := range names {
		fmt.Printf("names[%d]: %s\n", i, val)
	}
	fmt.Println()

	fmt.Println("distances\n====================")
	for i, val := range distances {
		fmt.Printf("distances[%d]: %d\n", i, val)
	}
	fmt.Println()

	fmt.Println("data\n====================")
	for i, val := range data {
		fmt.Printf("data[%d]: %d\n", i, val)
	}
	fmt.Println()

	fmt.Println("ratio\n====================")
	for i, val := range ratio {
		fmt.Printf("ratio[%d]: %.2f\n", i, val)
	}
	fmt.Println()

	fmt.Println("alives\n====================")
	for i, val := range alives {
		fmt.Printf("alives[%d]: %t\n", i, val)
	}
	fmt.Println()

	fmt.Println("zero\n====================")
	for i, val := range zero {
		fmt.Printf("zero[%d]: %d\n", i, val)
	}
	fmt.Println()

}
