package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	// Split array into slices.
	slicesNamesA := strings.Split(namesA, ", ")

	// Sort array namesA & namesB
	sort.Strings(slicesNamesA)
	sort.Strings(namesB)

	fmt.Printf("Names A: %v\n", slicesNamesA)
	fmt.Printf("Names B: %v\n", namesB)

	// Check each element if equal
	var isDifferent bool

	if len(slicesNamesA) == len(namesB) {
		for i := range slicesNamesA {
			if slicesNamesA[i] != namesB[i] {
				isDifferent = true
				break
			}
		}
	}

	if isDifferent {
		fmt.Println("They are NOT equal")
	} else {
		fmt.Println("They are equal")
	}
}
