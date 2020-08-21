// fmt package defines the fmt.Stringer interface: to allow any type to decide
// how it will be displayed when printed. It’s easy to set up any type to
// satisfy Stringer; just define a String() method that returns a string.

// Many functions in the fmt package check whether the values passed to
// them satisfy the Stringer interface, and call their String methods if so.
// This includes the Print, Println, and Printf functions and more.

package main

import "fmt"

type Gallons float64

func (g Gallons) String() string { // make gallons satisfy Stringr
	return fmt.Sprintf("%0.2f Gallons", g)
}

func main() {
	fmt.Println(Gallons(12.34535228906))
	fmt.Print(Gallons(345.134833), "\n")
	fmt.Printf("%v\n", Gallons(646.130687))
	fmt.Printf("%#v", Gallons(646.130687))
}
