package main

import (
	"fmt"
	"strconv"
)

func main() {
	// --- #1 ---
	//var games []string

	//games := []string{}
	//
	//fmt.Println("(before) length:", len(games))
	//fmt.Println("(before) cap:", cap(games))
	//
	//games = append(games, "pacman", "mario", "tetris", "doom")
	//fmt.Println("\n(after) length:", len(games))
	//fmt.Println("(after) cap:", cap(games))

	games := []string{"pacman", "mario", "tetris", "doom"}

	// --- #2 ---
	//for i := 0; i <= len(games); i++ {
	//	currentGame := games[:i]
	//	fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(currentGame), cap(currentGame))
	//}

	// --- #3 ---
	// it shows, zero will append and create new backing array due of exceeded previous capacity
	fmt.Println()

	zero := games[:0]
	fmt.Printf("games's len: %d cap: %d\n", len(games), cap(games))
	fmt.Printf("zeros's len: %d cap: %d\n", len(zero), cap(zero))

	for i := 0; i < 5; i++ {
		zero = append(zero, strconv.Itoa(i))
		fmt.Printf("zero's len: %d cap: %d\n", len(zero), cap(zero))
	}

	// --- #4 ---
	// It only iterate each element based on the length
	fmt.Println()
	for i := range zero {
		currentZero := zero[:i]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", i, len(currentZero), cap(currentZero))
	}

	// --- #5 ---
	// iterate over the capacity
	fmt.Println()
	zero = zero[:cap(zero)]
	for i, v := range zero {
		currentZero := zero[:i]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", i, len(currentZero), cap(currentZero), v)
	}

	// --- #6 ---
	// show different backing array.
	fmt.Println()
	zero[6] = "andri"
	games[3] = "boom"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", games)

	// --- #7 ---
	//newGames := games[:8]
	//fmt.Printf("new games : %q\n", newGames)
}
