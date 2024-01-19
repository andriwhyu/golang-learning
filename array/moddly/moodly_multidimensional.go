package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const usage = `[your name] [positive|negative]`

func main() {
	arguments := os.Args[1:]

	if len(arguments) != 2 {
		fmt.Println(usage)
		return
	}

	name, mood := arguments[0], arguments[1]

	moodIndex := 0
	moods := [...][3]string{
		{"Good 👍", "Awesome 😎", "Great ☺️"},
		{"Sad 😔", "Worst 😭", "Bad 👎"},
	}

	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := randomizer.Intn(len(moods[0]))

	if mood == "negative" {
		moodIndex = 1
	}

	fmt.Printf("%s feels %s\n", name, moods[moodIndex][randomIndex])

}
