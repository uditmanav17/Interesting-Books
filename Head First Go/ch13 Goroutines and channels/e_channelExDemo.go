package main

import (
	"fmt"
	"time"
)

func repoprtNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(time.Second)
	}
}

func send(myChannel chan string) {
	repoprtNap("sending routine", 2)
	fmt.Println("**sending value**")
	myChannel <- "a"
	fmt.Println("**sending value**")
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)

	go send(myChannel)
	repoprtNap("receiving channel", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
