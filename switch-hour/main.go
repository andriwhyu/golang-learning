package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch h := t.Hour(); {
	case h >= 6:
		fmt.Println("Morning")
	case h >= 12:
		fmt.Println("Afternoon")
	case h >= 18:
		fmt.Println("Evening")
	default:
		fmt.Println("Night")
	}
}
