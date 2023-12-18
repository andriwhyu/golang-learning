package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}

	scale := os.Args[1]

	// num, err := strconv.ParseFloat(os.Args[1], 64)

	// if err != nil {
	// 	fmt.Println("I couldn't get that, sorry.")
	// 	return
	// }

	switch scale {
	case "micro":
		fmt.Printf("%s's richter scale is less than 2.0\n", scale)
	case "very minor":
		fmt.Printf("%s's richter scale is 2 - 2.9\n", scale)
	case "minor":
		fmt.Printf("%s's richter scale is 3 - 3.9\n", scale)
	case "light":
		fmt.Printf("%s's richter scale is 4 - 4.9\n", scale)
	case "moderate":
		fmt.Printf("%s's richter scale is 5 - 5.9\n", scale)
	case "strong":
		fmt.Printf("%s's richter scale is 6 - 6.9\n", scale)
	case "major":
		fmt.Printf("%s's richter scale is 7 - 7.9\n", scale)
	case "great":
		fmt.Printf("%s's richter scale is 8 - 9.9\n", scale)
	case "massive":
		fmt.Printf("%s's richter scale is 10+\n", scale)
	default:
		fmt.Printf("%s's richter scale is unknown\n", scale)
	}
}
