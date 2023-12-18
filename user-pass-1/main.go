package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPass  = "Invalid password for %q.\n"
	accessed = "Access granted to %q.\n"

	u_data = "jack"
	p_data = "12345"
)

func main() {
	l := len(os.Args)

	if l != 3 {
		println(usage)
		return
	}

	u := os.Args[1]
	p := os.Args[2]

	if u != u_data {
		fmt.Printf(errUser, u)
	} else if p != p_data {
		fmt.Printf(errPass, u)
	} else {
		fmt.Printf(accessed, u)
	}

}
