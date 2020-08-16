package main

import (
	"fmt"
)

func sliceExample() {
	// slices are view of array
	underlyingArray := [7]string{"a", "b", "c", "d", "e"}
	slice1 := underlyingArray[0:3]
	fmt.Println(slice1)

	slice2 := underlyingArray[3:5]
	fmt.Println(slice2)
	slice2[0] = "X" // modify value in slice
	fmt.Println(slice2)
	fmt.Println(underlyingArray)
}

func sliceAppend() {
	slice := make([]string, 1) // need to use "make" to make an empty slice
	slice = append(slice, "b")
	fmt.Println(slice, len(slice))
	slice = append(slice, "c")
	fmt.Println(slice, len(slice))
	slice = append(slice, "d", "e")
	fmt.Println(slice, len(slice))
}

func defaultSliceVals() {
	// all slices defaults to nil
	var intSlice []int
	fmt.Printf("intSlice: %#v\n", intSlice)

	var strSlice []string
	fmt.Printf("strSlice: %#v\n", strSlice)

	var boolSlice []bool
	fmt.Printf("boolSlice: %#v\n", boolSlice)

	var floatSlice []float64
	fmt.Printf("floatSlice: %#v\n", floatSlice)

	fmt.Println(len(floatSlice))
}

func main() {
	// sliceExample()
	// sliceAppend()
	defaultSliceVals()
}
