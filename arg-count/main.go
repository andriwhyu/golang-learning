package main

import (
	"fmt"
	"os"

	"github.com/divan/num2words"
)

func main() {
	l := len(os.Args)
	message := "Give me args"

	// fmt.Println(num2words.Convert(17))
	if l > 1 {
		totalArg := l - 1
		sep := " "
		var argVal string

		for i := 1; i < l; i++ {
			if i == l-1 {
				sep = ""
			}

			argVal += os.Args[i] + sep
		}

		message = "There is " + num2words.Convert(totalArg) + ": " + "\"" + argVal + "\""
	}

	if l == 6 {
		message = "There are 5 arguments"
	}

	fmt.Println(message)
}
