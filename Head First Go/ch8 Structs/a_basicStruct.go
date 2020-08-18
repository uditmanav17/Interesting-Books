package main

import (
	"fmt"
)

var myStruct struct {
	number float64
	word   string
	toggle bool
}

func basicStruct() {
	fmt.Printf("%#v\n", myStruct)
	//assigning values to struct
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true
	fmt.Println("myStruct.word - ", myStruct.word)
	fmt.Printf("%#v\n", myStruct)
}

func main() {
	basicStruct()
}
