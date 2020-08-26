package main

import (
	"fmt"
)

func hi() {
	fmt.Println("Hi!")
}

func bye() {
	fmt.Println("Bye!")
}

func callTwice(myFunc func()) {
	myFunc()
	myFunc()
}

func main() {
	callTwice(hi)
	callTwice(bye)
}
