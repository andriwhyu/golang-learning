package main

import (
	"fmt"
	"os"
	"strconv"
)

const usage = "Please tell me numbers (maximum 5 numbers)."

func main() {
	args := os.Args[1:]

	var (
		num       [5]float64
		sum       float64
		totalData float64
	)

	if len(args) == 0 || len(args) > 5 {
		fmt.Println(usage)
		return
	}

	for i := range args {
		val, err := strconv.ParseFloat(args[i], 64)

		if err != nil {
			continue
		}
		num[i] = val
		sum += val
		totalData++
	}

	fmt.Printf("Your numbers: %v\n", num)
	fmt.Printf("Average: %v\n", sum/totalData)
}
