package main

import (
	"fmt"
	"strings"
	"time"
)

func printTime() {
	fmt.Println("Calling printTime")
	var now time.Time = time.Now()
	var year int = now.Year()
	month := now.Month()
	fmt.Println(now.Day(), month, year)
	fmt.Println()
}

func replaceChar() {
	fmt.Println("Calling replaceChar")
	broken := "G# R#cks!"
	replacer := strings.NewReplacer("#", "o")
	fmt.Println(replacer.Replace(broken))
}

func main() {
	// printTime()
	replaceChar()
}
