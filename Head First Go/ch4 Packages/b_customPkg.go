package main

import (
	"fmt"
	"greeting" // custom package
	"greeting/deutsch"
)

func main() {
	greeting.Hello()
	greeting.Hi()

	fmt.Println("\nAll Greetings: ")
	greeting.AllGreetings()

	fmt.Println("\nGerman Greetings: ")
	deutsch.Hallo()
	deutsch.GutenTag()
}
