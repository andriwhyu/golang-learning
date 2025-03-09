package main

import "fmt"

/*
	Notes:
	1. What I want to achieve in this code?
		- Simulate the lite process of dynamic array; append, insert, delete.
	2. Definition of success:
		- Able to run the code properly for at min 10 elements

	Boundaries:
	1. This code only support the simple process. It only supports ONE args per operation.
	2. Its only compatible for grade array.
	3. No empty array in the middle of array. It would be normalized by deleting the element.

	Leave the statement above. Update the challenge to create insert and delete with go append.
*/

func manualAppend(abc [3]int) {
	// calculate how many empty element to calculate

	// there are two cases of append: 1). there is an empty space in the array. 2). Array out of capacity.
}

func main() {
	studentGrade := [10]int{50, 40, 75, 88, 51, 77, 91, 89, -1, -1}
	_ = studentGrade

	manualAppend([3]int{})

	fmt.Println(-1e100)
}
