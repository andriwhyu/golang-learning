package main

import "fmt"

/*
unnamed type is type of data without specific naming.
Example:
[3]int{13, 14, 17} - unnamed types
[3]int - underlying types

Named type is an underlying types that named differently.
Example:
type bookcase [3]int

Unnamedtype is compareable with named type as long as it have the same underlying types.
Named type is not compareable with named type if the name is different. Even if the underlying type is the same.
*/

func main() {
	type (
		integer  int
		bookcase [5]int
		cabinet  [5]int

		accessories [5]integer
	)

	blue := bookcase{6, 7, 4, 3, 1}
	red := cabinet{6, 7, 4, 3, 1}
	//red := [5]int{6, 7, 4, 3, 1}

	_ = [5]integer{} == accessories{}
	//_ = [5]integer{} == bookcase{}

	fmt.Print("Are they the same? ")

	if blue == bookcase(red) {
		fmt.Println("Yay!")
	} else {
		fmt.Println("Boo!")
	}

	fmt.Printf("blue: %#v\n", blue)
	fmt.Printf("red: %#v\n", red)

	fmt.Println([5]integer{} == accessories{}) // should be false
	fmt.Println(blue == [5]int{6, 7, 4, 3, 1}) // should be false

	//if accessories == cabinet {
	//
	//}
}
