package main

import "fmt"

func stringArray() {
	var notes [7]string //array declaration

	notes[0] = "Udit"
	notes[1] = "Manav"
	notes[2] = "Udit"

	fmt.Println(notes)
	fmt.Printf("%#v\n", notes)

	for index, ele := range notes {
		fmt.Println(index, ele)
	}
}

func intArray() {
	var intArr [5]int = [5]int{9, 10, 0, 0, 26}

	intArr[0]++
	intArr[0]++
	intArr[3]++

	fmt.Println(intArr)
	fmt.Printf("%#v\n", intArr)

	for index, ele := range intArr {
		fmt.Println(index, ele)
	}
}

func main() {
	stringArray()
	intArray()
}
