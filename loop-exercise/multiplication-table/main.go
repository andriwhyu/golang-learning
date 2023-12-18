package main

import "fmt"

func horizontalLine() {
	var (
		line string
		i    = 1
	)

	for {
		if i > 35 {
			break
		}
		line += "-"
		i++
	}

	fmt.Println(line)
}

func main() {

	fmt.Printf("%5s%2s", "X", "|")

	for i := 0; i <= 5; i++ {
		if i == 0 {
			fmt.Printf("%3d", i)
			continue
		}
		fmt.Printf("%5d", i)
	}
	fmt.Println()
	horizontalLine()

	for i := 0; i <= 5; i++ {
		fmt.Printf("%5d%2s", i, "|")
		for j := 0; j <= 5; j++ {
			if j == 0 {
				fmt.Printf("%3d", i*j)
				continue
			}
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
