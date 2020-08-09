package main

import (
	"fmt"
	"reflect"
)

func pointerBasics() {
	a := 6
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(&a))
}

func pointerBasics2() {
	var myInt int
	var myIntPointer *int
	myInt = 5
	fmt.Println("Number:", myInt)

	myIntPointer = &myInt
	fmt.Println("Pointer", myIntPointer)

	myInt += 3
	fmt.Println("Number:", myInt)
	*myIntPointer += 2
	fmt.Println("Pointer:", *myIntPointer)
}

func negate(boolPointer *bool) {
	// fmt.Println(boolPointer)
	// fmt.Println(*boolPointer)
	*boolPointer = !*boolPointer
}

func boolModification() {
	boolean := true
	fmt.Println("Val before func call:", boolean)
	// fmt.Println(reflect.TypeOf(boolean))
	// fmt.Println(reflect.TypeOf(&boolean))
	negate(&boolean)
	fmt.Println("Val after func call:", boolean)
}

func main() {
	// pointerBasics()
	// pointerBasics2()
	boolModification()
}
