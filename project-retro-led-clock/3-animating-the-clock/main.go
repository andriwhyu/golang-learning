package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

type placeholder [5]string

func main() {
	// Simplify type definition

	separator := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	blankSeparator := placeholder{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	}

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

	screen.Clear()
	i := 1
	sep := separator

	for {
		now := time.Now()
		hour, mins, sec := now.Hour(), now.Minute(), now.Second()

		if i%2 == 0 {
			sep = blankSeparator
		}

		// [8][5]string
		clock := [...]placeholder{
			// extract the digits: 17 becomes, 1 and 7 respectively
			digits[hour/10], digits[hour%10],
			sep,
			digits[mins/10], digits[mins%10],
			sep,
			digits[sec/10], digits[sec%10],
		}

		screen.MoveTopLeft()
		for line := range clock[0] {
			for digit := range clock {
				fmt.Print(clock[digit][line], "  ")
			}
			fmt.Println()
		}
		sep = separator
		time.Sleep(1 * time.Second)
		i++
	}
}
