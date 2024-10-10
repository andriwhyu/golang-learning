package main

import (
	"fmt"
	s "github.com/inancgumus/prettyslice"
	"strings"
)

// It will return an index of the word where it found inside the slice.
// -1 means, the word is not exist in the slice.
func getIndex(slice []string, searchedWord string) int {
	for i, val := range slice {
		if val == searchedWord {
			return i
		}
	}
	return -1
}

func main() {
	// You need to add a newline after each sentence in another slice.
	// Don't touch the following code.
	lyric := strings.Fields(`yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday`)

	// ===================================
	//
	// ~~~ CHANGE THIS CODE ~~~

	// define the fix slice which len = len(lyric) + 3 because there will be 3 additional new line.
	fix := make([]string, len(lyric)+3)

	searchWords := []string{"away", "stay"}
	wordIndexes := make([]int, 0, len(searchWords)+1)

	// get the index of the searched word and store it in slice.
	for _, word := range searchWords {
		wordIndex := getIndex(lyric, word)
		wordIndexes = append(wordIndexes, wordIndex)
	}

	// add last element of slice as part of the word indexes to add new line in the end of sentence.
	wordIndexes = append(wordIndexes, len(lyric)-1)

	var (
		startIndexLyric int
		startIndexFix   int
	)

	// count the number of line that will have been added. Init value is 1.
	countNewLine := 1

	for _, wordIndex := range wordIndexes {
		// get the range of lyric before the new line appended.
		sliceLyric := lyric[startIndexLyric : wordIndex+1]
		startIndexLyric = wordIndex + 1

		endIndexFix := wordIndex + countNewLine
		// copy the lyric to the new slice and add new line after it.
		copy(fix[startIndexFix:endIndexFix], sliceLyric)
		copy(fix[endIndexFix:endIndexFix+1], []string{"\n"})

		countNewLine++
		startIndexFix = wordIndex + countNewLine
	}

	// ===================================

	// another approach here:
	// https://github.com/inancgumus/learngo/blob/master/16-slices/exercises/25-add-lines/solution/main.go

	// Currently, it prints every sentence on the same line.
	// Don't touch the following code.
	s.Show("fix slice", fix)

	for _, w := range fix {
		fmt.Print(w)
		if w != "\n" {
			fmt.Print(" ")
		}
	}
}

func init() {
	//
	// YOU DON'T NEED TO TOUCH THIS.
	//
	// This initializes some options for the prettyslice package.
	// You can change the options if you want.
	//
	// This code runs before the main function above.
	//
	// s.Colors(false)     // if your editor is light background color then enable this
	//
	s.PrintBacking = true  // prints the backing arrays
	s.MaxPerLine = 5       // prints max 15 elements per line
	s.SpaceCharacter = '⏎' // print this instead of printing a newline (for debugging)
}
