package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	argsLen := len(os.Args)

	if argsLen < 2 {
		fmt.Println("> Enter your name...")
		return
	}
	mood := [...]string{"sad 😔", "terrible 😥", "happy ☺️", "awesome 😎"}
	name := os.Args[1]

	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))

	i := randomizer.Intn(len(mood))
	fmt.Printf("%s feels %s\n", name, mood[i])

	//for j := 0; j < 5; j++ {
	//	i := randomizer.Intn(len(mood))
	//	fmt.Printf("%s feels %s\n", name, mood[i])
	//}

}
