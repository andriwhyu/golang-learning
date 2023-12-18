// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

func main() {
	l_args := len(os.Args)

	if l_args == 1 {
		fmt.Println("Give me a year number")
		return
	} else if l_args > 2 {
		fmt.Println("Too many arguments")
		return
	}

	arg := os.Args[1]
	year, err := strconv.Atoi(arg)

	if err != nil {
		fmt.Printf("%q is not a valid year.\n", arg)
		return
	}

	var message string

	if year%100 == 0 {
		message = fmt.Sprintf("%d is not a leap year.", year)
	} else if year%4 == 0 || year%400 == 0 {
		message = fmt.Sprintf("%d is a leap year.", year)
	}

	fmt.Println(message)
}
