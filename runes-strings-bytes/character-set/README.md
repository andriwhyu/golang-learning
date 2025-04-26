# README

The idea of the program is simulated how byte, strings, and runes work.

# Bytes & Strings

Underlying of strings is a slice of bytes. It mean, string can be converted into a slice byte (`[]byte`).

Otherwise, a slice byte can be converted into string.

```go
// convert byte to string
str := string([]byte{101, 123, 150}

// cobvert string to byte
byteStr := []byte(str)
```

String is immutable data type. It mean a value of a string can’t be replaced. If we re-assign a string in existing variable,  under the hood, it create a new memory. With this condition, we can’t replace the character of string.

```go
strString := "hey"
strString[0] = "m" // it will return an error
```

Conversion from string to byte will create a new byte backing array. Which mean generate a byte from string and adjust the byte will not change the existing byte but change the new created byte.

```go
strString := "hey"
byteStr := []byte(strString)
byteStr[0] = 'm'

fmt.Println(strString) // it will print `"hey"`
fmt.Println(string(byteStr)) // it will print `"mey"`
```

Conversion from byte to string and replace it to the prior variable will not replace the backing array but actually it create new backing array under the hood.

```go
strString := "hey"
byteStr := []byte(strString)
byteStr[0] = 'm'

strString = string(byteStr) // it not change the backing array, but it pointing to the new backing array
```

# Runes

There are two term in runes, rune literal and runes. **Rune literal** is a type less data type, which mean a single character of rune can be represent in any kind of integer; `uint`, `uint32`, `uint64`.

**Runes** is a code point of rune literal. Runes is not equal with byte, because a single character of rune can be store in multiple byte when it encoded in UTF-8. By default, Go will encode a string literal into UTF-8.

To calculate a length of character, especially multibyte character, used `utf8.RuneCountInString(str)` function instead of `len()` . `len()` calculate number of byte, meanwhile the `utf8.RuneCountInString(str)` function calculate per character len.

```go
// calculate string length
multiByteStr := "Euro(£)"
fmt.Println("standard len count:", len(multiByteStr)) // result: 8 with 7 char

// calculate rune/character length
fmt.Println("rune count of string:", utf8.RuneCountInString(multiByteStr)) // result: 7 with 7 char
```

If we’ve a string, we can loop over each character and see the indexes and print the memory byte.

When loop over each character of string, noticed that indexing of character is not based on character number but it followed the index where the rune started.

```go
customStr := "euro £ ❤️"

for i, val := range customStr {
	fmt.Printf("customStr[%-2d] = % -8x = %q\n", i, string(val), val)
}

/* result
customStr[0 ] = 65       = 'e'
customStr[1 ] = 75       = 'u'
customStr[2 ] = 72       = 'r'
customStr[3 ] = 6f       = 'o'
customStr[4 ] = 20       = ' '
customStr[5 ] = c2 a3    = '£'   -> 2 bytes, index from 5 and jump to 7 for next character
customStr[7 ] = 20       = ' '
customStr[8 ] = e2 9d a4 = '❤'
customStr[11] = ef b8 8f = '️'
*/
```

## Work with Rune

Working with multibyte byte characters also can be proceed with `[]rune` slice. A string convert into slice of rune and can be indexing directly even if it’s multibyte. It more simple on access the index or character.

```go
// comparison
customStr := "euro £ ❤️"
customRune := []rune(customStr)

for i, val := range customRune {
	fmt.Printf("customRune[%-2d] = % -8x = %q\n", i, string(val), val)
}

/* Result

customRune[0 ] = 65       = 'e'
customRune[1 ] = 75       = 'u'
customRune[2 ] = 72       = 'r'
customRune[3 ] = 6f       = 'o'
customRune[4 ] = 20       = ' '
customRune[5 ] = c2 a3    = '£'   -> No jumping index
customRune[6 ] = 20       = ' '
customRune[7 ] = e2 9d a4 = '❤'   -> No jumping index
customRune[8 ] = ef b8 8f = '️'
*/

// -----------------------

// print single character
fmt.Printf("print from byte slice: %c\n", customStr[5])
fmt.Printf("print from byte slice: %s\n", customStr[5:7])
fmt.Printf("print from rune slice: %c\n", customRune[5])

/* Result

print from byte slice: Â   -> print incorrect character
print from byte slice: £   -> print correct character but required a slice range and string conversion
print from rune slice: £   -> print correct character without slice range and string convertion

*/
```

**Why most developer more prefer to use []byte instead of []rune when working with character?**

1. Conversion from string to []rune is resource consumption. Due it’s not recommend to use it especially if the character was changing dynamically and we will often to convert a string to a rune slices.
2. []byte is most common for developer, it reflect by a lot of package or library that support byte operation instead of rune slice operation.