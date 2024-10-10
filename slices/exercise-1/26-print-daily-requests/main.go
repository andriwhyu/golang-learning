package main

import (
	"fmt"
	"strings"
)

func main() {
	// There are 3 request totals per day (8-hourly)
	const N = 3

	// DAILY REQUESTS DATA (8-HOURLY TOTALS PER DAY)
	reqs := []int{
		500, 600, 250, // 1st day: 1350 requests
		200, 400, 50, // 2nd day: 650 requests
		900, 800, 600, // 3rd day: 2300 requests
		750, 250, 100, // 4th day: 1100 requests
		// grand total: 5400 requests
	}

	// ================================================
	// #1: Make a new slice with the exact size needed.

	// calculate len(req) % 3. If it's 0 then size = len(req)/N
	// otherwise size = len(req)/N + 1

	lenReqs := len(reqs)
	size := lenReqs / N

	if lenReqs%N != 0 {
		size++
	}

	daily := make([][]int, 0, size)

	// ================================================

	// ================================================
	// #2: Group the `reqs` per day into the slice: `daily`.
	//
	// So the daily will be:
	// [
	//  [500, 600, 250]
	//  [200, 400, 50]
	//  [900, 800, 600]
	//  [750, 250, 100]
	// ]

	// 1st approach

	//dailyData := make([]int, 0, N)
	//for i, val := range reqs {
	//	dailyData = append(dailyData, val)
	//
	//	if (i+1)%N == 0 || i+1 == lenReqs {
	//		daily = append(daily, dailyData)
	//		dailyData = make([]int, 0, N)
	//	}
	//}

	// 2nd approach
	for N <= len(reqs) {
		daily = append(daily, reqs[:N]) // append the daily requests
		reqs = reqs[N:]                 // move the slice pointer for the next day
	}

	// add the residual data
	if len(reqs) > 0 {
		daily = append(daily, reqs)
	}

	// ================================================
	// #3: Print the results

	// Print a header
	fmt.Printf("%-10s%-10s\n", "Day", "Requests")
	fmt.Println(strings.Repeat("=", 20))

	// Loop over the daily slice and its inner slices to find
	// the daily totals and the grand total.
	// ...
	var grandTotal int
	for i, dailySlice := range daily {
		var totalReq int

		for _, req := range dailySlice {
			totalReq += req

			fmt.Printf("%-10d%-10d\n", i+1, req)
		}

		fmt.Printf("%9s %-10d\n\n", "TOTAL:", totalReq)

		grandTotal += totalReq
	}
	fmt.Printf("%9s %-10d\n", "GRAND:", grandTotal)

	// ================================================
}
