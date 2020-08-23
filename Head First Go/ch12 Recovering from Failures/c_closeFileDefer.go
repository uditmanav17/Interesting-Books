package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(filename string) (*os.File, error) {
	fmt.Println("Opening", filename)
	return os.Open(filename)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

func GetFloats2(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName) // replaced w
	if err != nil {
		return nil, err
	}
	defer CloseFile(file) // moved just below open file, will exec even if error occurs

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil

}

func main() {
	// filepath := "./dataError.txt"
	filepath := "./data.txt"
	numbers, err := GetFloats2(filepath)
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum %0.2f", sum)
}
