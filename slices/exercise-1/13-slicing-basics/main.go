package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := "2 4 6 1 3 5"

	numsStr := strings.Split(data, " ")

	var (
		evens []int
		odds  []int
		nums  []int
	)

	for _, e := range numsStr {
		eInt, err := strconv.Atoi(e)

		if err != nil {
			fmt.Println("Invalid input")
			return
		}

		nums = append(nums, eInt)

		if eInt%2 == 0 {
			evens = append(evens, eInt)
			continue
		}

		odds = append(odds, eInt)
	}

	numsLength := len(nums)
	middleIndex := numsLength / 2

	middle := nums[middleIndex-1 : middleIndex+1]
	first2 := nums[:2]
	last2 := nums[numsLength-2:]

	evensLast1 := evens[len(evens)-1:]
	oddsLast2 := odds[len(odds)-2:]

	fmt.Printf("nums                : %v\n", nums)
	fmt.Printf("evens               : %v\n", evens)
	fmt.Printf("odds                : %v\n", odds)
	fmt.Printf("middle              : %v\n", middle)
	fmt.Printf("first 2             : %v\n", first2)
	fmt.Printf("last 2              : %v\n", last2)
	fmt.Printf("evens last 1        : %v\n", evensLast1)
	fmt.Printf("odds last 2         : %v\n", oddsLast2)
}
