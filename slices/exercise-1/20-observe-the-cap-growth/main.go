package main

import "fmt"

const iterate = 1e7

func main() {
	var growthSlice []int

	prevCap := cap(growthSlice)

	for i := 0; i < iterate; i++ {
		growthSlice = append(growthSlice, i)
		curCap := cap(growthSlice)

		if curCap > prevCap {
			fmt.Printf("len:%-15d cap:%-15d growth:%.2f\n", len(growthSlice), curCap, float64(curCap)/float64(prevCap))
		}
		prevCap = curCap
	}
}
