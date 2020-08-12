package main

import (
	"fmt"
	"greeting" // custom package
	"greeting/deutsch"
	"keyboard"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	fmt.Println("\nAll Greetings: ")
	greeting.AllGreetings()
	fmt.Println("\nGerman Greetings: ")
	deutsch.Hallo()
	deutsch.GutenTag()
	// to freeze screen when creating exe
	inp, err := keyboard.GetFloat()
	fmt.Println(inp, err)
}
