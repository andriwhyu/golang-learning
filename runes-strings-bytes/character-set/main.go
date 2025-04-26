package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	// convert string to byte and vice versa.
	say := "🔥hello"
	byteSay := []byte(say)
	fmt.Println(byteSay)
	fmt.Println(string(byteSay))
	fmt.Println(string(byteSay[0])) // it will return unintended character due emoji required more than 1 byte
	fmt.Println()

	var charStart, charStop int
	args := os.Args[1:]

	if len(args) < 2 || len(args) > 2 {
		fmt.Println("Invalid input, input should ONLY contains 2 params")
		return
	}

	charStart, _ = strconv.Atoi(args[0])
	charStop, _ = strconv.Atoi(args[1])

	if charStart == 0 || charStop == 0 {
		charStart = 'A'
		charStop = 'Z'
	}

	fmt.Printf("%-10s %-10s %-10s %-10s\n%s\n", "char", "dec", "hex", "encoded", strings.Repeat("-", 45))

	for i := charStart; i <= charStop; i++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -10x\n", i, string(i))
	}

	var byteTest byte = 'À' // acceptable due it's a single byte (192)
	_ = byteTest

	//var byteTestMultiple byte = '✅' // unacceptable due it's a multiple byte (192)
	//_ = byteTest

	// acceptable due it's rune type
	runeTest := '✅'
	_ = runeTest
	fmt.Println()

	// calculate string length
	multiByteStr := "Euro(£)"
	fmt.Println("standard len count:", len(multiByteStr)) // result: 8 with 7 char

	// calculate rune/character length
	fmt.Println("rune count of string:", utf8.RuneCountInString(multiByteStr)) // result: 7 with 7 char

	// calculate rune len from byte
	multiByte := []byte(multiByteStr)
	fmt.Println("rune count of byte:", utf8.RuneCount(multiByte)) // result: 7 with 7 char

	fmt.Println()

	// loop over character of multibyte string
	customStr := "euro £ ❤️"
	for i, val := range customStr {
		fmt.Printf("customStr[%-2d] = % -8x = %q\n", i, string(val), val)
	}

	fmt.Println()

	// comparison working with []rune and []slice
	customRune := []rune(customStr)
	for i, val := range customRune {
		fmt.Printf("customRune[%-2d] = % -8x = %q\n", i, string(val), val)
	}

	fmt.Println()

	// print single character
	fmt.Printf("print from byte slice: %c\n", customStr[5])
	fmt.Printf("print from byte slice: %s\n", customStr[5:7])
	fmt.Printf("print from rune slice: %c\n", customRune[5])
}
