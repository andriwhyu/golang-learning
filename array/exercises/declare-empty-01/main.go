package main

import "fmt"

func main() {
	var (
		names     [3]string
		distances [5]int
		data      [5]uint
		ratio     [1]float64
		alives    [4]bool
		zero      [0]uint8
	)

	fmt.Printf("names%9s %#v\n", ":", names)
	fmt.Printf("distances%5s %#v\n", ":", distances)
	fmt.Printf("data%10s %#v\n", ":", data)
	fmt.Printf("ratio%9s %#v\n", ":", ratio)
	fmt.Printf("alives%8s %#v\n", ":", alives)
	fmt.Printf("zero%10s %#v\n\n", ":", zero)

	fmt.Printf("names%9s %T\n", ":", names)
	fmt.Printf("distances%5s %T\n", ":", distances)
	fmt.Printf("data%10s %T\n", ":", data)
	fmt.Printf("ratio%9s %T\n", ":", ratio)
	fmt.Printf("alives%8s %T\n", ":", alives)
	fmt.Printf("zero%10s %T\n\n", ":", zero)

	fmt.Printf("names%9s %q\n", ":", names)
	fmt.Printf("distances%5s %d\n", ":", distances)
	fmt.Printf("data%10s %d\n", ":", data)
	fmt.Printf("ratio%9s %.2f\n", ":", ratio)
	fmt.Printf("alives%8s %t\n", ":", alives)
	fmt.Printf("zero%10s %d\n", ":", zero)

}
