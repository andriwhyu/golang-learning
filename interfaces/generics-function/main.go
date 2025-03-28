package main

import "fmt"

// this project will demonstrate how interface can handle all data types
// and process interface value in specific data type process.
// also this project will explain how to use generics to simplify specific operation data type.

func main() {
	// called doAnythingSwitch
	fmt.Println("doAnythingSwitch...")
	doAnythingSwitch(1)
	doAnythingSwitch(2.4)
	doAnythingSwitch("Hi")
	doAnythingSwitch(true)

	// called doAnything
	fmt.Println("\ndoAnything...")
	doAnything(1)
	doAnything(2.4)
	doAnything("Hi")
	doAnything(true)

	// called add(a, b) function
	// the limitation of this approach is the value will return as any. if it's so we need
	// to convert it first to apply the additional operation
	fmt.Println("\nadd...")
	intResult := add(1, 2)
	if intResult != nil {
		fmt.Println(intResult)
	}

	floatResult := add(1.2, 2.5)
	if floatResult != nil {
		fmt.Println(floatResult)
	}

	// called addGenerics function
	fmt.Println("\naddGenerics...")

	intGenerics := addGenerics(1, 2)
	fmt.Println(intGenerics)

	floatGenerics := addGenerics(1.2, 2.5)
	fmt.Println(floatGenerics)

}

// doAnythingSwitch received an interface parameter and print the value by read
// the data type using switch and val.(type)
func doAnythingSwitch(val any) {
	switch val.(type) {
	case int:
		fmt.Println("Integer:", val)
	case float64:
		fmt.Println("Float:", val)
	case string:
		fmt.Println("String:", val)
	default:
		fmt.Println("Unknown type:", val)
	}
}

// doAnything received an interface parameter and print the value by read the data type
// using val.(datatype) cast.
func doAnything(val interface{}) {
	intVal, ok := val.(int)
	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := val.(float64)
	if ok {
		fmt.Println("Float:", floatVal)
		return
	}

	stringVal, ok := val.(string)
	if ok {
		fmt.Println("String:", stringVal)
		return
	}

	fmt.Println("Unknown type:", val)
}

// add will apply "+" operation of a and b parameters (any) data type.
// it will return an any value.
func add(a, b any) any {
	aIntVal, aOkInt := a.(int)
	bIntVal, bOkInt := b.(int)

	if aOkInt && bOkInt {
		return aIntVal + bIntVal
	}

	aFloatVal, aOkFloat := a.(float64)
	bFloatVal, bOkFloat := b.(float64)

	if aOkFloat && bOkFloat {
		return aFloatVal + bFloatVal
	}

	aStringVal, aOkString := a.(string)
	bStringVal, bOkString := b.(string)

	if aOkString && bOkString {
		return aStringVal + " " + bStringVal
	}

	return nil
}

// addGenerics implement the generic placeholder. with generic, the return value will follow
// the parameter type automatically. it happens because the generics will set the value when it was called.
func addGenerics[T int | float64 | string](a, b T) T {
	return a + b
}
