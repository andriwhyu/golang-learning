package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	EUR = iota
	GBP
	JPY
)

const usage = `

Usage:
go run main.go [number]
go run main.go 10.0
`

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Print("Please provide the amount to be converted.", usage)
		return
	}

	usd, err := strconv.ParseFloat(args[0], 64)

	if err != nil {
		fmt.Print("Invalid amount. It should be a number.", usage)
		return
	}

	convertRate := [...]float64{
		EUR: 0.88,
		GBP: 0.78,
		JPY: 113.02,
	}

	fmt.Printf("%.2f USD is %.2f EUR\n", usd, usd*convertRate[EUR])
	fmt.Printf("%.2f USD is %.2f GBP\n", usd, usd*convertRate[GBP])
	fmt.Printf("%.2f USD is %.2f JPY\n", usd, usd*convertRate[JPY])
}
