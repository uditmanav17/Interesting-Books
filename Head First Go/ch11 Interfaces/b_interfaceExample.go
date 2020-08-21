package main

import (
	"fmt"
	"mypkg"
)

func main() {
	var value mypkg.MyInterface // declare an interface
	value = mypkg.MyType(5)
	value.MethodWithParams(5)
	value.MethodWithoutParams()
	// value.MethodNotInInterface()  // can't use this
	fmt.Println(value.MethodWithReturn())
}
