package main

import (
	"fmt"
	"github.com/mattn/go-runewidth"
)

//const creationSLA = time.Second * 600

//const (
//	defaultGopher = "pocket"
//	usage         = "the variety of gopher"
//)

//type User struct {
//	Name  string `validate:"required"`       // This field is required
//	Email string `validate:"required,email"` // This field is required and must be a valid email
//}

func main() {
	//words := strings.Fields("hi hello hola")
	//words1 := "hi hello hola"
	//
	//for j := 0; j < len(words); j++ {
	//	fmt.Printf("#%-2d: %q\n", j+1, words[j])
	//}
	//
	//for _, v := range words1 {
	//	fmt.Println(v)
	//}
	//
	//fmt.Printf("creation SLA: %s\n", creationSLA)

	//var (
	//	gopherType string
	//	gopherAge  int
	//)
	//
	//flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	//flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")
	//flag.IntVar(&gopherAge, "g_age", 0, "int of gopher age")
	//
	//flag.Usage = func() {
	//	flag.CommandLine.SetOutput(os.Stdout)
	//	_, err := fmt.Fprintf(flag.CommandLine.Output(), "Kegunaan dari apps %s:\n", os.Args[0])
	//	if err != nil {
	//		return
	//	}
	//
	//	flag.PrintDefaults()
	//}
	//
	//flag.Parse()
	//
	//fmt.Println(gopherType)
	//fmt.Println(gopherAge)

	//validate := validator.New()
	//
	//// Create an instance of User with invalid data
	//user := User{
	//	Name:  "",                     // This will trigger a validation error
	//	Email: "invalid-email-format", // This will also trigger a validation error
	//}
	//
	//// Validate the struct
	//err := validate.Struct(user)
	//if err != nil {
	//	// If there are validation errors, print them
	//	var validationErrors validator.ValidationErrors
	//	if errors.As(err, &validationErrors) {
	//		for _, validationError := range validationErrors {
	//			fmt.Printf("Field %s failed validation with tag: %s\n", validationError.Field(), validationError.Tag())
	//		}
	//	}
	//} else {
	//	fmt.Println("Validation passed!")
	//}
	//prettyslice.MaxPerLine = 10
	//prettyslice.PrintBacking = true
	//
	//randomAlphabet := "b c f g d e"
	//randomAlphabetSlice := strings.Fields(randomAlphabet)
	//
	//correctAlphabet := append([]string{"a"}, randomAlphabetSlice...)
	//prettyslice.Show("1st iteration", correctAlphabet)
	//correctAlphabet = append(correctAlphabet, correctAlphabet[3:5]...)
	//prettyslice.Show("2nd iteration", correctAlphabet)
	//correctAlphabet = append(correctAlphabet[:3], correctAlphabet[5:]...)
	//prettyslice.Show("3rd iteration", correctAlphabet)
	//
	//fmt.Printf("%#v", correctAlphabet[:cap(correctAlphabet)])

	//s := []int{1, 2, 3, 4, 5}
	//
	//prettyslice.Show("All", s)
	//prettyslice.Show("s[2:]", s[2:])
	//prettyslice.Show("s[0:3]", s[0:3])
	//
	//copy(s[2:], s[0:3]) // Copy first three elements to positions starting at index 2
	//fmt.Println(s)

	//go func() {
	//	for {
	//		fmt.Println("Hi it's me, go func")
	//	}
	//}()
	//
	//fmt.Println("Hi, it's me, non go func")
	//time.Sleep(2 * time.Second)

	// Example JSON string
	//jsonData := `{"name": "Alice", "age": 25, "active": true}`

	// Query the JSON
	//result := gjson.Get(jsonData, "name")
	//fmt.Println(result)

	//ch := make(chan string) // Unbuffered channel
	//
	//// Goroutine for sending data
	//go func() {
	//	ch <- "Order Received"
	//	fmt.Println("Sent: Order Received")
	//
	//	ch <- "Order Processed"
	//	fmt.Println("Sent: Order Processed")
	//}()
	//
	//// Simulate some processing time in the receiver
	//time.Sleep(2 * time.Second)
	//
	//// Receive data from the channel
	//fmt.Println("Received:", <-ch)
	//time.Sleep(1 * time.Second) // Simulate processing time
	//fmt.Println("Received:", <-ch)

	//a := []byte{1, 2, 3}
	//b := []byte{10}
	//
	//copy(a[:], b)
	//
	//fmt.Println(a)
	//width := 5
	//height := 3
	//runeSlice := make([]rune, 0, width*height)
	//
	//prettyslice.Show("runeSlice before:", runeSlice)
	//runeSlice = append(runeSlice, '\n')
	//prettyslice.Show("runeSlice after:", runeSlice)

	fmt.Println(len("☆"))
	fmt.Println(runewidth.StringWidth("つ"))

}

//func init() {
//	prettyslice.PrintBacking = true
//	prettyslice.MaxPerLine = 20
//}
