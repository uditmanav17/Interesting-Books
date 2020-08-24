package main

import (
	"fmt"
)

func greeting(myChannel chan string, s string) {
	myChannel <- s
}

func example1() {
	fmt.Println("\nExample 1 -")
	myChannel := make(chan string)
	go greeting(myChannel, "hi")
	go greeting(myChannel, "bye")

	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func example2() {
	fmt.Println("\nExample 2 -")
	channel1 := make(chan string)
	channel2 := make(chan string)

	go abc(channel1)
	go def(channel2)

	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
}

func main() {
	// example1()
	example2()
}
