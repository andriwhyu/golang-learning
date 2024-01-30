package main

import (
	"fmt"
	"strings"
)

const (
	firstname = iota
	lastname
	nickname
)

func main() {
	scientistNames := [...][3]string{
		{"Albert", "Eistein", "time"},
		{"isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	fmt.Printf("%s\t%s\t%s\n", "First Name", "Last Name", "Nickname")
	fmt.Println(strings.Repeat("=", 50))

	for _, person := range scientistNames {
		fmt.Printf("%s\t\t%s\t\t%s\n", person[firstname], person[lastname], person[nickname])
	}

}
