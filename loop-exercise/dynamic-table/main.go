package main

import (
	"fmt"
	"os"
	"strconv"
)

func horizontalLine(dimension int) {
	var (
		line string
		i    = 1
	)

	for {
		if i > (dimension+2)*5 {
			break
		}
		line += "-"
		i++
	}

	fmt.Println(line)
}

func main() {

	args := os.Args

	if len(args) != 2 {
		fmt.Println("Unexpected input")
		return
	}

	dimension, err := strconv.Atoi(args[1])

	if err != nil || dimension <= 0 {
		fmt.Println("Invalid input")
		return
	}

	fmt.Printf("%5s%2s", "X", "|")

	for i := 0; i <= dimension; i++ {
		if i == 0 {
			fmt.Printf("%3d", i)
			continue
		}
		fmt.Printf("%5d", i)
	}
	fmt.Println()
	horizontalLine(dimension)

	for i := 0; i <= dimension; i++ {
		fmt.Printf("%5d%2s", i, "|")
		for j := 0; j <= dimension; j++ {
			if j == 0 {
				fmt.Printf("%3d", i*j)
				continue
			}
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
