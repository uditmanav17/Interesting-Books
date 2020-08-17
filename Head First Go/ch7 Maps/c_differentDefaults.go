package main

import (
	"fmt"
)

func status(name string) {
	grades := map[string]int{"Alice": 90, "bob": 50}
	if grades[name] < 60 {
		fmt.Println(name, "is Failing!")
	} else {
		fmt.Println(name, "passed!")
	}
}

func mapDifferentFromDefault() {
	status("Alice")
	status("Carl") // shows carl is failing but carl isn't registered in grades
}

func status2(name string) {
	grades := map[string]int{"Alice": 90, "bob": 50}
	grade, value := grades[name]
	if value {
		if grade < 60 {
			fmt.Println(name, "is Failing!")
		} else {
			fmt.Println(name, "passed!")
		}
	} else {
		fmt.Println(name, "didn't appear for exam!")
	}

}
func mapSafeAccess() {
	status2("Alice")
	status2("Carl")
}

func mapDeleting() {
	var isPrime map[int]bool     //declares map var
	isPrime = make(map[int]bool) // actually creates a map

	isPrime[5] = true
	_, ok := isPrime[5]
	fmt.Println("Is 5 in isPrime?", ok)

	delete(isPrime, 5)
	_, ok = isPrime[5]
	fmt.Println("Is 5 in isPrime?", ok)
}

func main() {
	fmt.Println("\nExecuting mapDifferentFromDefault:")
	mapDifferentFromDefault()

	fmt.Println("\nExecuting mapSafeAccess:")
	mapSafeAccess()

	fmt.Println("\nExecuting mapDeleting:")
	mapDeleting()
}
