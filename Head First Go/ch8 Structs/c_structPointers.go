package main

import (
	"fmt"
)

type testStruct struct {
	number float64
	word   string
	toggle bool
}

func updateNum(s testStruct, n float64) {
	s.number = n
	fmt.Println("In Function - %#v", s)
}

func updateNum2(s *testStruct, n float64) {
	s.number = n
	fmt.Println("In Function - %#v", s)
}

func main() {
	var myStruct testStruct
	fmt.Printf("%#v\n", myStruct)
	//assigning values to struct
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true
	fmt.Println("myStruct.word - ", myStruct.word)
	fmt.Printf("%#v\n", myStruct)

	// updateNum to update number
	updateNum(myStruct, 3.14159)
	fmt.Printf("Returned - %#v\n", myStruct) // doesn't update, because call by value

	// updateNum2 to update number
	updateNum2(&myStruct, 3.14159)
	fmt.Printf("Returned - %#v\n", myStruct) // updates, because call by reference
}
