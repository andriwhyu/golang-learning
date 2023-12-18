package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}

	num, err := strconv.ParseFloat(os.Args[1], 64)

	if err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	}

	switch {
	case num >= 10.0:
		fmt.Println("massive")
	case num >= 8.0:
		fmt.Println("great")
	case num >= 7.0:
		fmt.Println("major")
	case num >= 6.0:
		fmt.Println("strong")
	case num >= 5.0:
		fmt.Println("moderate")
	case num >= 4.0:
		fmt.Println("light")
	case num >= 3.0:
		fmt.Println("minor")
	case num >= 2.0:
		fmt.Println("very minor")
	default:
		fmt.Println("micro")
	}
}
