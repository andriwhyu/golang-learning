package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	usage   = "Please tell me numbers (maximum 5 numbers)."
	caution = "Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers..."
)

func swap(a, b float64) (float64, float64) {
	return b, a
}

func bubleSort(arr [5]float64) [5]float64 {
	numLength := len(arr)

	for i := 0; i < numLength-1; i++ {
		for j := 0; j < numLength-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = swap(arr[j], arr[j+1])
			}
		}
	}

	return arr
}

func main() {
	args := os.Args[1:]

	var (
		num [5]float64
	)

	if len(args) == 0 {
		fmt.Println(usage)
		return
	} else if len(args) > 5 {
		fmt.Println(caution)
		return
	}

	for i := range args {
		val, err := strconv.ParseFloat(args[i], 64)

		if err != nil {
			continue
		}
		num[i] = val
	}

	num = bubleSort(num)

	fmt.Printf("Your numbers: %v\n", num)
}
