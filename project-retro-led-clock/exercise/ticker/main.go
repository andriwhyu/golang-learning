package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

const totalDigit = 8

func main() {
	sep := separator
	adjustor := 1

	var index int

	for {
		screen.Clear()
		screen.MoveTopLeft()
		now := time.Now()
		hour, mins, sec := now.Hour(), now.Minute(), now.Second()

		if sec%2 == 0 {
			sep = blankSeparator
		}

		clock := [...]placeholder{
			// extract the digits: 17 becomes, 1 and 7 respectively
			digits[hour/10], digits[hour%10],
			sep,
			digits[mins/10], digits[mins%10],
			sep,
			digits[sec/10], digits[sec%10],
		}

		if index >= totalDigit {
			adjustor = -1
			index += adjustor
		}

		if index < 0 {
			adjustor = 1
			index += adjustor
		}

		// Create fade out effect
		if adjustor > 0 {
			if clock[index] == separator {
				index += adjustor
			}

			for line := range clock[0] {
				for digit := range clock[index:] {
					fmt.Print(clock[digit][line], "  ")
				}
				fmt.Println()
			}
		}

		// Create fade in effect
		if adjustor < 0 {
			for line := range clock[0] {
				for i := range clock[:totalDigit] {
					if i >= index {
						fmt.Print(clock[i-index][line], "  ")
					} else {
						fmt.Print(blankSeparator[line], "  ")
					}
				}
				fmt.Println()
			}
		}

		sep = separator
		index += adjustor
		time.Sleep(1 * time.Second)
	}
}

//func main() {
//	sep := ":"
//	adjustor := 1
//
//	var index int
//
//	for {
//		screen.Clear()
//		screen.MoveTopLeft()
//		now := time.Now()
//		hour, mins, sec := now.Hour(), now.Minute(), now.Second()
//		hour1, hour2 := strconv.Itoa(hour/10), strconv.Itoa(hour%10)
//		mins1, mins2 := strconv.Itoa(mins/10), strconv.Itoa(mins%10)
//		sec1, sec2 := strconv.Itoa(sec/10), strconv.Itoa(sec%10)
//
//		clock := [...]string{hour1, hour2, sep, mins1, mins2, sep, sec1, sec2}
//
//		if index >= totalDigit {
//			adjustor = -1
//			index += adjustor
//		}
//
//		if index < 0 {
//			adjustor = 1
//			index += adjustor
//		}
//
//		if adjustor > 0 {
//			if clock[index] == ":" {
//				index += adjustor
//			}
//
//			for _, s := range clock[index:] {
//				fmt.Print(s)
//			}
//		}
//
//		if adjustor < 0 {
//			for i, _ := range clock[:totalDigit] {
//				if i >= index {
//					fmt.Print(clock[i-index])
//				} else {
//					fmt.Print(" ")
//				}
//			}
//		}
//
//		index += adjustor
//		time.Sleep(1 * time.Second)
//	}
//}
