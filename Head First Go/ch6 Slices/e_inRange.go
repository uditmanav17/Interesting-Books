package main

import (
	"fmt"
)

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	fmt.Println(inRange(0, 100, 45.4, 78.5, 123.7, 789.0, 4.0, -32.5))
	fmt.Println(inRange(-10, 10, 4.1, -12, 12, 5.3, -4.5))
}
