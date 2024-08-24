package main

import (
	"fmt"
	"sort"
)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)
	midIndex := len(items) / 2
	midSlice := items[midIndex-1 : midIndex+2]
	// sorting a slice will also change the backing array which reflect the main slice also got sorted
	sort.Strings(midSlice)
	fmt.Println("Sorted  :", items)
}
