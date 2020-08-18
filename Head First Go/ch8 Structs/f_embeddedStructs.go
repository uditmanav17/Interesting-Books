package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func embeddedStructExample1() {
	// make an Address struct then add it to Employee
	address := magazine.Address{Street: "123 Oak St", City: "Delhi", PostalCode: "112200"}
	emp1 := magazine.Employee{Name: "Udit"}
	emp1.HomeAddress = address
	fmt.Printf("%#v\n", emp1)
}

func embeddedStructExample2() {
	// add Address by chaining
	emp1 := magazine.Employee{Name: "Manav"}
	emp1.HomeAddress.Street = "456 Oak St"
	emp1.HomeAddress.City = "Noida"
	fmt.Printf("%#v\n", emp1)
}

func anonymouStruct() {
	emp1 := magazine.Employee2{Name: "Anonymous"}
	emp1.Address.Street = "--REDACTED--"
	emp1.Address.City = "--REDACTED--"

	// advantage of anon struct
	emp1.PostalCode = "--REDACTED--" // no need to write emp1.Address.PostalCode
	fmt.Printf("%#v\n", emp1)
}

func main() {
	embeddedStructExample1()
	embeddedStructExample2()
	anonymouStruct()
}
