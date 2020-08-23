package main

import (
	"fmt"
	"log"
)

func find(item string, items []string) bool {
	for _, sliceItem := range items {
		if sliceItem == item {
			return true
		}
	}
	return false
}

type Refrigerator []string

func (r Refrigerator) open() {
	fmt.Println("Opening Refrigerator")
}

func (r Refrigerator) close() {
	fmt.Println("Closing Refrigerator\n")
}

func (r Refrigerator) findFood(food string) error {
	r.open()
	defer r.close()
	if find(food, r) {
		fmt.Println("Found", food)
	} else {
		return fmt.Errorf("%s not in Refrigerator", food)
	}
	return nil
}

func main() {
	fridge := Refrigerator{"milk", "butter", "banana"}
	for _, food := range []string{"milk", "fruit"} {
		fmt.Println("Searching for", food)
		err := fridge.findFood(food)
		if err != nil {
			log.Fatal(err)
		}
	}
}
