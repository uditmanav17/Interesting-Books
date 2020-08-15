package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt") //open data file for reading
	// if error in opening file, report and exit
	if err != nil {
		log.Fatal(err)
	}
	// create a new scanner for file
	scanner := bufio.NewScanner(file)
	// loop unitl EOF is reached and scanner.Scan returns False
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// if error in closing file, report and exit
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// if error in in scanning file, report and exit
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}
