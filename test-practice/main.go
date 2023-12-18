package main

import "fmt"

func main() {
	// words := strings.Fields("hi hello hola")
	words1 := "hi hello hola"

	// for j := 0; j < len(words); j++ {
	// 	fmt.Printf("#%-2d: %q\n", j+1, words[j])
	// }

	for _, v := range words1 {
		fmt.Println(v)
	}
}
