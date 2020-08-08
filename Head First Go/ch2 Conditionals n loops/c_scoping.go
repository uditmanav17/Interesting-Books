package main

import (
	"fmt"
)

var packageVar = "PackageVar"

func main() {
	var functionVar = "Function"
	if true {
		var conditionalVar = "Conditional"
		fmt.Println(packageVar)
		fmt.Println(functionVar)
		fmt.Println(conditionalVar)
	}
	fmt.Println(packageVar)
	fmt.Println(functionVar)
	fmt.Println(conditionalVar) // out of scope
}
