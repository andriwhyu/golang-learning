package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1st approach
	//var letter_type int
	//
	//cons := "bcdfghjklmnpqrstvxz"
	//vowel := "aiueo"
	//semi_vowel := "wy"
	//
	//l := len(os.Args)
	//message := "Give me a letter"

	//if l == 2 && len(os.Args[1]) == 1 {
	//	str_args := os.Args[1]
	//
	//	letter_type = strings.IndexAny(cons, str_args)
	//
	//	if letter_type > -1 {
	//		message = "\"" + str_args + "\"" + " is a consonant."
	//	}
	//
	//	letter_type = strings.IndexAny(vowel, str_args)
	//
	//	if letter_type > -1 {
	//		message = "\"" + str_args + "\"" + " is a vowel."
	//	}
	//
	//	letter_type = strings.IndexAny(semi_vowel, str_args)
	//
	//	if letter_type > -1 {
	//		message = "\"" + str_args + "\"" + " is sometimes a vowel, sometimes not."
	//	}
	//}
	//
	//print(message)

	// 2nd apporach
	/*
		The improvement:
		- Not necessary to check both 3 conditions, we only need to check 2, meanwhile the last one keep it as default.
		- Change the variable naming, using camel case instead of snack case.
	*/

	const (
		cons      = "bcdfghjklmnpqrstvxz"
		semiVowel = "wy"
	)
	args := os.Args[1:]
	lenArgs := len(args)
	msg := "Give me a letter"

	if lenArgs == 1 && len(args[0]) == 1 {
		letter := args[0]

		msg = "is a vowel."
		letterIndex := strings.IndexAny(cons, letter)

		if letterIndex > -1 {
			msg = "is a consonant."
		}

		letterIndex = strings.IndexAny(semiVowel, letter)

		if letterIndex > -1 {
			msg = "is sometimes a vowel, sometimes not."
		}

		msg = fmt.Sprintf("%q %s", letter, msg)
	}

	fmt.Println(msg)

}
