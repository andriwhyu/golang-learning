package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	usage               = "Provide only the [starting] and [stopping] positions"
	errInvalidPosition  = "Wrong positions. Starting index >= 0 && stopping pos < %d\n"
	errInvalidInput     = "Invalid input"
	errStartMoreThanEnd = "Wrong positions. Starting index < Stopping index"
)

func main() {
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}
	dataLength := len(ships)

	args := os.Args[1:]

	if len(args) == 0 || len(args) > 2 {
		fmt.Println(usage)
		return
	}

	start, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println(errInvalidInput)
		fmt.Println(err)
		return
	}

	end := dataLength

	if len(args) > 1 {
		end, err = strconv.Atoi(args[1])

		if err != nil {
			fmt.Println(errInvalidInput)
			fmt.Println(err)
			return
		}
	}

	if start < 0 || end > dataLength {
		fmt.Printf(errInvalidPosition, dataLength)
		return
	}

	if start > end {
		fmt.Println(errStartMoreThanEnd)
		return
	}
	fmt.Printf("%v\n", ships[start:end])
}
