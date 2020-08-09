package main

import (
	"fmt"
)

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Printf("%2d times: %s\n", i, line)
	}
}

func main() {
	repeatLine("hello", 5)
}
