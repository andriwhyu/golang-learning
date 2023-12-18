package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		n := rand.Intn(10)

		fmt.Println(n)
	}
}
