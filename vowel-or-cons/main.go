package main

import (
	"os"
	"strings"
)

func main() {
	var letter_type int

	cons := "bcdfghjklmnpqrstvxz"
	vowel := "aiueo"
	semi_vowel := "wy"

	l := len(os.Args)
	message := "Give me a letter"

	if l == 2 && len(os.Args[1]) == 1 {
		str_args := os.Args[1]

		letter_type = strings.IndexAny(cons, str_args)

		if letter_type > -1 {
			message = "\"" + str_args + "\"" + " is a consonant."
		}

		letter_type = strings.IndexAny(vowel, str_args)

		if letter_type > -1 {
			message = "\"" + str_args + "\"" + " is a vowel."
		}

		letter_type = strings.IndexAny(semi_vowel, str_args)

		if letter_type > -1 {
			message = "\"" + str_args + "\"" + " is sometimes a vowel, sometimes not."
		}
	}

	print(message)

}
