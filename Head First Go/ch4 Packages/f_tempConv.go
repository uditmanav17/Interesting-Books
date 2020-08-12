package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/keyboard"
)

func main() {
	fmt.Println("Enter Temp in F: ")
	fTemp, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fTemp - 32) * 5 / 9
	fmt.Printf("%0.2f degrees in Celsius", celsius)
}
