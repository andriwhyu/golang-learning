package main

import (
	"fmt"
	"time"
)

type placeholder [5]string

//func printDigit(digits [10]placeholder, numbers string, index int) {
//	for j := 0; j < len(numbers); j++ {
//		digit, _ := strconv.Atoi(string(numbers[j]))
//		fmt.Print(digits[digit][index], "  ")
//	}
//}

func main() {
	// Simplify type definition

	separator := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
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

	//timeNow := time.Now()
	//hour, minutes, seconds := strconv.Itoa(timeNow.Hour()), strconv.Itoa(timeNow.Minute()), strconv.Itoa(timeNow.Second())
	//
	//if len(hour) < 2 {
	//	hour = fmt.Sprintf("0%s", hour)
	//}
	//
	//if len(minutes) < 2 {
	//	minutes = fmt.Sprintf("0%s", minutes)
	//}
	//
	//if len(seconds) < 2 {
	//	seconds = fmt.Sprintf("0%s", seconds)
	//}
	//
	//for i := range digits[0] {
	//	printDigit(digits, hour, i)
	//	fmt.Print(separator[i], "  ")
	//	printDigit(digits, minutes, i)
	//	fmt.Print(separator[i], "  ")
	//	printDigit(digits, seconds, i)
	//	fmt.Println()
	//}

	now := time.Now()
	hour, mins, sec := now.Hour(), now.Minute(), now.Second()
	fmt.Println(0/10, 0%10)

	fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, mins, sec)

	// [8][5]string
	clock := [...]placeholder{
		// extract the digits: 17 becomes, 1 and 7 respectively
		digits[hour/10], digits[hour%10],
		separator,
		digits[mins/10], digits[mins%10],
		separator,
		digits[sec/10], digits[sec%10],
	}

	for line := range clock[0] {
		for digit := range clock {
			fmt.Print(clock[digit][line], "  ")
		}
		fmt.Println()
	}
}
