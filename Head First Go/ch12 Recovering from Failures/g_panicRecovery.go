package main

import (
	"fmt"
)

func calmDown() {
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	}
}

func main() {
	defer calmDown()
	err := fmt.Errorf("an error occured")
	panic(err)
}
