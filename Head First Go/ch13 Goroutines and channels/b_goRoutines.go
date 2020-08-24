package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
	// fmt.Println()
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
	// fmt.Println()
}

func simpleMain() {
	fmt.Println("Running simpleMain...")
	a()
	b()
	fmt.Println("\nCompleted simpleMain\n\n")
}

func goMain() {
	fmt.Println("Running simpleMain...")
	go a()
	go b()
	go a()
	go b()
	go a()
	go b()

	time.Sleep(time.Second)
	fmt.Println("\nCompleted simpleMain\n\n")
}

func main() {
	simpleMain()

	goMain()
}
