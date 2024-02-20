package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Simplify type definition
	type placeholder [5]string

	digits_0 := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	digits_1 := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	digits_2 := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	digits_3 := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	digits_4 := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	digits_5 := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	digits_6 := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	digits_7 := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	digits_8 := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	digits_9 := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := [...]placeholder{
		digits_0,
		digits_1,
		digits_2,
		digits_3,
		digits_4,
		digits_5,
		digits_6,
		digits_7,
		digits_8,
		digits_9,
	}

	//for line := range digits[0] {
	//	// Print a line for each placeholder in digits
	//	for digit := range digits {
	//		fmt.Print(digits[digit][line], "  ")
	//	}
	//	fmt.Println()
	//}

	printedDigits := "023"

	_, err := strconv.Atoi(printedDigits)

	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	for i := range digits[0] {
		for j := 0; j < len(printedDigits); j++ {
			digit, _ := strconv.Atoi(string(printedDigits[j]))
			fmt.Print(digits[digit][i], "  ")
		}
		fmt.Println()
	}
}
