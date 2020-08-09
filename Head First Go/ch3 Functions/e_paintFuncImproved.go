// adding error Handling
package main

import (
	"fmt"
	"log"
)

var metersPerL float64

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 || height < 0 {
		return 0, fmt.Errorf("width/height can't be negative")
	}
	area := width * height
	return area / metersPerL, nil
}

func main() {
	metersPerL = 1
	paint, err := paintNeeded(5, -3.3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2fL paint Needed", paint)

}
