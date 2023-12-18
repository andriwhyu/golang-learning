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
)

func main() {
	u_data := [2]string{"jack", "inanc"}
	p_data := [2]string{"1888", "1879"}

	var message string

	l := len(os.Args)

	if l != 3 {
		println(usage)
		return
	}

	u := os.Args[1]
	p := os.Args[2]

	for index, user := range u_data {
		if u != user {
			message = fmt.Sprintf(errUser, u)
			continue
		} else if p != p_data[index] {
			message = fmt.Sprintf(errPass, u)
		} else {
			message = fmt.Sprintf(accessed, u)
			break
		}
	}

	fmt.Print(message)

}
