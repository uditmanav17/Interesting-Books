package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input, err)

	if grade >= 100 {
		status := "Perfect!"
	} else if grade >= 60 {
		status := "Pass!"
	} else {
		status := "Fail!"
	}

	fmt.Println("A grade of ", grade, " is ", status)
}
