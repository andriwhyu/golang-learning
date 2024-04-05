package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

const (
	nanoSecondDevider = 1e8
)

func main() {
	screen.Clear()
	sep := separator
	dotSep := dot

	for {
		now := time.Now()
		hour, mins, sec, nanoSec := now.Hour(), now.Minute(), now.Second(), now.Nanosecond()

		if sec%2 == 0 {
			sep, dotSep = blankSeparator, blankSeparator
		}

		// [8][5]string
		clock := [...]placeholder{
			// extract the digits: 17 becomes, 1 and 7 respectively
			digits[hour/10], digits[hour%10],
			sep,
			digits[mins/10], digits[mins%10],
			sep,
			digits[sec/10], digits[sec%10],
			dotSep,
			digits[nanoSec/nanoSecondDevider],
		}

		screen.MoveTopLeft()
		for line := range clock[0] {
			for digit := range clock {
				fmt.Print(clock[digit][line], "  ")
			}
			fmt.Println()
		}
		sep, dotSep = separator, dot
		time.Sleep(100 * time.Millisecond)
	}
}
