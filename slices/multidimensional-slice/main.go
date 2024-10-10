package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	Use case:
	- We want to calculate total spent per day base on the data.
	- We will use multidimensional slices. Each element will contain slices data represent a day spend.
		- Example:
			{
				{100, 230, 1000}, // 1st day
				{300, 100, 150},  // 2nd day
			}
*/

type spendingType [][]int

func main() {
	// Testing data
	//spendingData := [][]int{
	//	{100, 230, 1000},
	//	{50, 10},
	//	{400},
	//	{300, 100, 150},
	//}

	// empty array definition
	//var spendingData [][]int
	//spendingData = append(spendingData, []int{200, 300, 5})
	//spendingData = append(spendingData, []int{30, 120})

	spendingData, err := fetchData()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Total spending")
	for i, daySpend := range spendingData {
		var sumSpend int

		for _, val := range daySpend {
			sumSpend += val
		}

		fmt.Printf("Day %d: %v\n", i+1, sumSpend)
	}
}

func fetchData() (spendingType, error) {
	spendingRaw := `100 230 1000
50 10
400
300 100 150`

	// split the data based on new line
	spendingSlice := strings.Split(spendingRaw, "\n")
	// init the returned slice
	spendingDataResult := make(spendingType, 0, len(spendingSlice))

	for _, daySpendStr := range spendingSlice {
		// split each spend based on whitespace
		eachSpend := strings.Fields(daySpendStr)
		// init the converted slice
		daySpendSlice := make([]int, 0, len(eachSpend))

		for _, valStr := range eachSpend {
			spendInt, err := strconv.Atoi(valStr)
			if err != nil {
				return nil, fmt.Errorf("cannot convert value, msg: %w", err)
			}

			daySpendSlice = append(daySpendSlice, spendInt)
		}

		spendingDataResult = append(spendingDataResult, daySpendSlice)
	}
	return spendingDataResult, nil
}
