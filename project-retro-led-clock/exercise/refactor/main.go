package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

func main() {
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
