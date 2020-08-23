package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	filePath := "./data.txt"
	numbers, err := datafile.GetFloats2(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0.0
	for _, num := range numbers {
		sum += num
	}
	arrLen := float64(len(numbers))
	fmt.Println("Sum - ", sum)
	fmt.Println("Average - ", sum/arrLen)

}
