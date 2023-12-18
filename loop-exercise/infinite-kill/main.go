package main

import (
	"fmt"
	"time"
)

func main() {
	// String
	charset := `/-\|`
	i := 0

	for {
		fmt.Printf("\r%s Please Wait. Processing.......", string(charset[i]))
		time.Sleep(250 * time.Millisecond)
		i++
		if i > len(charset)-1 {
			i = 0
		}
	}
}
