package main

import (
	"fmt"
)

func simpleMap() {
	myMap := map[string]float64{"a": 23., "b": 76.}
	for k, v := range myMap {
		fmt.Println(k, v)
	}
}

func simpleMap2() {
	var isPrime map[int]bool     //declares map var
	isPrime = make(map[int]bool) // actually creates a map

	isPrime[5] = true
	isPrime[6] = false

	for k, v := range isPrime {
		fmt.Println(k, v)
	}
}

func mapDefaults() {
	numbers := map[string]int{}

	numbers["Assigned"] = 12
	fmt.Printf("Assigned - %#v\n", numbers["Assigned"])
	fmt.Printf("Not Assigned - %#v\n", numbers["Not Assigned"]) // val defaults to 0

	for k, v := range numbers {
		fmt.Println(k, v)
	}

	words := make(map[string]string)
	words["Assigned"] = "YEAH"
	fmt.Printf("Assigned - %#v\n", words["Assigned"])
	fmt.Printf("Not Assigned - %#v\n", words["Not Assigned"]) // val defaults to ""
}

func main() {
	fmt.Println("Executing Simple Map:")
	simpleMap()

	fmt.Println("Executing Simple Map2:")
	simpleMap2()

	fmt.Println("Executing Map Defaults:")
	mapDefaults()
}
