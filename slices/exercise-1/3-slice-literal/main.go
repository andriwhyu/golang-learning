package main

import "fmt"

func main() {
	names := []string{"Syamil", "Gozy", "Nasution"}
	distances := []int{5, 2, 3, 1, 4}
	data := []uint8{9, 3, 2, 55, 6}
	ratios := []float64{1400.4, 150.32}
	alives := []bool{true, false, false, true}

	fmt.Printf("names     : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances : %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data      : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios    : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives    : %T %d %t\n", alives, len(alives), alives == nil)

	if len(distances) == len(data) {
		fmt.Println("The length of the distances and the data slices are the same.")
	}
}
