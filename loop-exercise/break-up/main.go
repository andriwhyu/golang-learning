package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	usage       = "Usage: [Min int] [Max int]. Min < Max"
	invalid_err = "Invalid input"
)

func main() {
	args := os.Args[1:]

	var (
		sum int
		sep = "+"
	)

	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	minVal, err_min := strconv.Atoi(args[0])
	maxVal, err_max := strconv.Atoi(args[1])

	if err_min != nil || err_max != nil || minVal > maxVal {
		fmt.Println(invalid_err)
		fmt.Println(usage)
		return
	}

	if maxVal%2 == 1 {
		maxVal -= 1
	}

	i := minVal

	for {
		if i%2 == 0 {
			if i == maxVal {
				sep = "="
			}

			fmt.Printf("%d %s ", i, sep)
			sum += i
		}

		i++
		if i > maxVal {
			break
		}
	}

	fmt.Printf("%d \n", sum)
}
