package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func horizontalLine(dimension int) {
	var (
		line string
		i    = 1
	)

	for {
		if i > (dimension+2)*5 {
			break
		}
		line += "-"
		i++
	}

	fmt.Println(line)
}

func calculation(ops string, num1, num2 int) string {
	var (
		num1_float = float32(num1)
		num2_float = float32(num2)
		result     string
	)

	switch ops {
	case `/`:
		if num2_float == 0 {
			num1_float = 0
			num2_float = 1
		}
		result = fmt.Sprintf("%.1f", num1_float/num2_float)
	case "+":
		result = fmt.Sprintf("%d", num1+num2)
	case "-":
		result = fmt.Sprintf("%d", num1-num2)
	case "%":
		if num2 == 0 {
			num1 = num1 + 1
			num2 = num1
		}
		result = fmt.Sprintf("%d", num1%num2)
	default:
		result = fmt.Sprintf("%d", num1*num2)
	}

	return result
}

const (
	ops = `*/+-%`
)

func main() {
	args := os.Args

	var (
		usage         = fmt.Sprintf("Usage: [op=%s] [size]", ops)
		error_message string
	)

	switch len(args) {
	case 2:
		error_message = "Size is missing"
		fmt.Printf("%s\n%s\n", error_message, usage)
		return
	case 3:
		check_ops := strings.IndexAny(ops, args[1])
		if check_ops == -1 {
			error_message = "Invalid operator."
			usage = fmt.Sprintf("Valid ops one of: %s", ops)

			fmt.Printf("%s\n%s\n", error_message, usage)
			return
		}
	default:
		fmt.Println(usage)
		return
	}

	dimension, err := strconv.Atoi(args[2])
	current_ops := args[1]

	if err != nil || dimension <= 0 {
		fmt.Println("Invalid input")
		return
	}

	// Logical stuff

	fmt.Printf("%5s%2s", current_ops, "|")

	for i := 0; i <= dimension; i++ {
		if i == 0 {
			fmt.Printf("%3d", i)
			continue
		}
		fmt.Printf("%5d", i)
	}
	fmt.Println()
	horizontalLine(dimension)

	for i := 0; i <= dimension; i++ {
		fmt.Printf("%5d%2s", i, "|")
		for j := 0; j <= dimension; j++ {
			if j == 0 {
				fmt.Printf("%3s", calculation(current_ops, i, j))
				continue
			}
			fmt.Printf("%5s", calculation(current_ops, i, j))
		}
		fmt.Println()
	}
}
