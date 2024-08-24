package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch h := t.Hour(); {
	case h >= 6 && h < 12:
		fmt.Println("Morning")
		break
	case h >= 12 && h < 18:
		fmt.Println("Afternoon")
		break
	case h >= 18 && h < 20:
		fmt.Println("Evening")
		break
	default:
		fmt.Println("Night")
	}

	fmt.Println(t.Hour())
}
