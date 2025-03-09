package main

import "fmt"

func main() {
	age := 32
	agePointers := &age

	fmt.Println(adultAges(agePointers))
	fmt.Println("before mutating:", *agePointers)
	adultAgesMutate(agePointers)
	fmt.Println("after mutating:", *agePointers)
}

func adultAges(age *int) int {
	return *age - 18
}

// mutate variable directly
func adultAgesMutate(age *int) {
	*age = *age - 18
}
