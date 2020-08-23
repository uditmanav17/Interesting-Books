package main

import (
	"fmt"
	"log"
)

func defFunc() {
	// f we add the defer keyword, then that call won’t be run until
	// all the remaining code in the function runs
	defer fmt.Println("Goodbye!")
	fmt.Println("Hi!")
	fmt.Println("Nice Weather, eh?")
}

func normalFunc() {
	fmt.Println("Goodbye!")
	fmt.Println("Hi!")
	fmt.Println("Nice Weather, eh?")
}

func defFuncError() error {
	defer fmt.Println("Goodbye!")
	fmt.Println("Hi!")
	return fmt.Errorf("I don't wanna talk.")
	fmt.Println("Nice Weather, eh?")
	return nil
}

func main() {
	fmt.Println("Normal Func -")
	normalFunc()
	fmt.Println("\nDefered Func -")
	defFunc()
	fmt.Println("\nDefered Func Errored -")
	err := defFuncError()
	if err != nil {
		log.Fatal(err)
	}
}
