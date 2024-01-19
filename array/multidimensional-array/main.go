package main

import "fmt"

func main() {
	students := [...][3]float64{
		{6, 9, 10},
		{5, 6, 8},
	}

	var sum float64

	for _, grades := range students {
		for _, grade := range grades {
			sum += grade
		}
	}

	totalData := float64(len(students) * len(students[0]))

	fmt.Printf("Average score: %.2f\n", sum/totalData)
}
