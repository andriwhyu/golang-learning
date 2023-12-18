package main

import "fmt"

func main() {
	var (
		sum int
		sep = "+"
	)

	for i := 1; i <= 10; i++ {
		if i == 10 {
			sep = " ="
		}
		fmt.Printf("%-2d%-2s", i, sep)
		sum += i
	}

	fmt.Printf("%3d\n", sum)
}
