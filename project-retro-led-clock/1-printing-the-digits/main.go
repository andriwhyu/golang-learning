package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Simplify type definition
	type placeholder [5]string

	digits0 := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	digits1 := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	digits2 := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	digits3 := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	digits4 := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	digits5 := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	digits6 := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	digits7 := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	digits8 := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	digits9 := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := [...]placeholder{
		digits0,
		digits1,
		digits2,
		digits3,
		digits4,
		digits5,
		digits6,
		digits7,
		digits8,
		digits9,
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
