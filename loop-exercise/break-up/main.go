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

	min, err_min := strconv.Atoi(args[0])
	max, err_max := strconv.Atoi(args[1])

	if err_min != nil || err_max != nil || min > max {
		fmt.Println(invalid_err)
		fmt.Println(usage)
		return
	}

	for {
		if min > max {
			break
		} else if min%2 == 1 {
			continue
		}

		if min == max {
			sep = "="
		}
		fmt.Printf("%d %s ", min, sep)
		sum += min
		min++

	}

	fmt.Printf("%d \n", sum)
}
