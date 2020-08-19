package main

import (
	"fmt"
)

type Number int

func (num *Number) Double() {
	*num *= 2
}

func main() {
	n := Number(4)
	fmt.Println("Number -", n)
	n.Double()
	fmt.Println("Number -", n)

	// TIP: This won't work because to call a method that requires a pointer receiver,
	// you have to be able to get a pointer to the value!
	// fmt.Println(Number(3).Double())
}
