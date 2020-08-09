package main

import (
	"fmt"
)

func basicFormatting() {
	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 15)
	fmt.Printf("A string: %s\n", "Manav")
	fmt.Printf("A bool: %t\n", true)
	fmt.Printf("values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n")
}

func widthFormatting() {
	// formatting with widths
	fmt.Printf("%f Liters needed\n", 1.819999999)

	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents") // 12 width for 1st arg, none for 2nd
	fmt.Println("----------------------------------")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5) // Padding of 1
	fmt.Printf("%12s | %2d\n", "Tape", 199)      // adjusted to accomodate 3 digits
}

func fractionalFormatting() {
	fmt.Printf("%%7.3f: %7.3f\n", 12.3456) // rounded to 3 places
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456) // rounded to 2 places
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456) // rounded to 1 places
	fmt.Printf("%%.1f: %.1f\n", 12.3456)   // rounded to 1 places, no padding
	fmt.Printf("%%.2f: %.2f\n", 12.3456)   // rounded to 2 places, no padding

	fmt.Printf("%.2f\n", 1.26000000000002)
	fmt.Printf("%.2f\n", 1.21999999999998)
}

func main() {
	// basicFormatting()
	// widthFormatting()
	fractionalFormatting()
}
