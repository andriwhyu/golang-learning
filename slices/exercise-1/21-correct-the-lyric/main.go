package main

import (
	"fmt"
	"strings"
)

func main() {
	// DON'T TOUCH THIS:
	lyric := strings.Fields(`all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay`)

	// ADD YOUR CODE BELOW:
	// 1st approach
	//var correctLyric []string
	//
	//sliceLyric1 := lyric[11:12]
	//sliceLyric2 := lyric[:7]
	//sliceLyric3 := lyric[7:12]
	//sliceLyric4 := lyric[12:]
	//
	//correctLyric = append(correctLyric, sliceLyric1...) // new backing array
	//correctLyric = append(correctLyric, sliceLyric2...) // new backing array
	//correctLyric = append(correctLyric, sliceLyric4...)
	//correctLyric = append(correctLyric, sliceLyric3...) // new backing array
	//
	//fmt.Printf("%s\n", correctLyric)

	// 2nd approach
	lyric = append([]string{"yesterday"}, lyric...)
	lyric = append(lyric, lyric[8:13]...)
	lyric = append(lyric[:8], lyric[13:]...)

	fmt.Printf("%s\n", lyric)
}
