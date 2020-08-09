package main

import (
	"fmt"
)

var metersPerL float64

// no return function
func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2fL paint needed\n", area)
}

// function returning float
func paintNeeded2(width float64, height float64) float64 {
	area := width * height
	return area / metersPerL
}

func main() {
	metersPerL = 1
	paintNeeded(4.2, 3.0)
	paintNeeded(5.2, 3.5)
	paintNeeded(5, 3.3)

	fmt.Printf("%.2fL\n", paintNeeded2(4.2, 3.0))
	fmt.Printf("%.2fL\n", paintNeeded2(5.2, 3.5))
	fmt.Printf("%.2fL\n", paintNeeded2(5, 3.3))

}
