package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func main() {
	fmt.Println(maximum(71.8, 56.4))
	fmt.Println(maximum(12, 567, 345, 3569, 123))
}
