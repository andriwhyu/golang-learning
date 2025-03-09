package main

import "fmt"

type HundlerFunc func(int, int) int

func (f HundlerFunc) MyFunction(num1, num2 int) int {
	return f(num1, num2)
}

func home(x, y int) int {
	return x + y
}

func main() {
	handler := HundlerFunc(home)
	fmt.Println(handler.MyFunction(2, 3))
}
