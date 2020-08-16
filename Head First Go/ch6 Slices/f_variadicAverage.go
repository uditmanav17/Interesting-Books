// Package average2 calculates the average of several numbers
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func main() {
	arguments := os.Args[1:]

	var numbers []float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	// average function expects one or more flat64 types, not slice
	fmt.Println("Average - ", average(numbers...)) // to pass numbers slice we need ... at end
}

// usage: go run f_variadicAverage.go 12 23 45
