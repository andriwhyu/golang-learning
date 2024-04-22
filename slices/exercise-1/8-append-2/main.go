package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		pizza       []string
		departures  []time.Time
		graduations []int
		lights      []bool
	)

	pizza = append(pizza, "papperoni", "roasted beef", "extra cheese")
	departures = append(departures, time.Now(), time.Now().Add(time.Hour*12), time.Now().Add(time.Hour*24))
	graduations = append(graduations, 1990, 2000, 1892)
	lights = append(lights, false, true, true)

	fmt.Printf("pizza\t\t: %v\n", pizza)
	fmt.Printf("departures\t: %v\n", departures)
	fmt.Printf("graduations\t: %v\n", graduations)
	fmt.Printf("lights\t\t: %v\n", lights)
}
