package main

import (
	"fmt"
)

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

func printPartInfo(p part) {
	fmt.Println("Description -", p.description)
	fmt.Println("Count -", p.count)
}

// assigning values to structs
func assign2Struct() {
	var porche car // define new var of type car
	porche.name = "Porche 911 R"
	porche.topSpeed = 300
	fmt.Printf("%#v\n", porche)

	var bolts part
	bolts.description = "Hex Bolts"
	bolts.count = 1500
	fmt.Printf("%#v\n", bolts)
	printPartInfo(bolts)
}

func main() {
	assign2Struct()
}
