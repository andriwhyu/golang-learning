package main

import "fmt"

func main() {
	array1, array2 := make([]string, 0), make([]string, 0)
	array1 = []string{"a", "b", "c", "x"}
	array2 = []string{"z", "y", "i"}
	var isCommon bool

main:
	for _, s := range array1 {
		for _, s2 := range array2 {
			if s == s2 {
				isCommon = true
				break main
			}
		}
	}

	fmt.Println(isCommon)
}
