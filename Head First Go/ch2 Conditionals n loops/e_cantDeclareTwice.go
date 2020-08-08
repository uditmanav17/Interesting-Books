// Only one variable in a short variable declaration has to be new

package main

import "fmt"

func declaringTwice() {
	a := 1
	// a := 2
	fmt.Println(a)
}

func declaringTwiceShort() {
	a := 1
	b, a := 2, 3
	c, b := 4, 5
	fmt.Println(a, b, c)
}

func main() {
	// declaringTwice()
	declaringTwiceShort()
}
