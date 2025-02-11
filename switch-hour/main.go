package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// in go, switch-case not necessarily to add break in the end of case. Go will automatically add it explicitly.
	switch h := t.Hour(); {
	case h >= 6 && h < 12:
		fmt.Println("Morning")
	case h >= 12 && h < 18:
		fmt.Println("Afternoon")
	case h >= 18 && h < 20:
		fmt.Println("Evening")
	default:
		fmt.Println("Night")
	}

	fmt.Println(t.Hour())
}
