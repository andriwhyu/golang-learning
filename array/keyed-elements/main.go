package main

import "fmt"

/*
One of key elements function is increase readability of code.
It can initialize array in specific index and left other index with default value 0 or ""
*/

const (
	// Iota only works for constant. The init value is 0. Will be increasing by one from top to bottom
	ETH = 9 - iota
	WAN
	ICX
)

func main() {
	array1 := [3]int{
		1: 123,
	}

	array2 := [...]int{
		4: 17,
		15,
		0: 12,
	}

	fmt.Printf("%v\n", array1)
	fmt.Printf("%v\n", array2)
	fmt.Println()

	//	Simple bitcoin convertion

	ratio := [...]float64{
		ETH: 25.5,
		WAN: 120.5,
		ICX: 120,
	}

	fmt.Printf("1 BTC is %g ETH\n", ratio[ETH])
	fmt.Printf("1 BTC is %g WAN\n", ratio[WAN])
	fmt.Printf("1 BTC is %g ICX\n", ratio[ICX])

	fmt.Printf("All data: %v", ratio)

}
