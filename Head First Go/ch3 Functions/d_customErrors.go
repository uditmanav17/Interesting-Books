package main

import (
	"errors"
	"fmt"
	"log"
)

func errorLog() {
	err := errors.New("Height can't be negative")
	fmt.Println(err)
	log.Fatal(err)
}

func manyReturns() (int, float64, bool) {
	return 1, 2.3333, true
}

func main() {
	fmt.Println(manyReturns())
	errorLog()
}
