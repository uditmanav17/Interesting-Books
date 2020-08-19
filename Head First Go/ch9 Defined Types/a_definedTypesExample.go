package main

import (
	"fmt"
)

func main() {
	type Liters float64  // defined type
	type Gallons float64 // defined type

	busFuel := Liters(10.5)  // convert float64 in Liters, busFuel is now type Liters
	carFuel := Gallons(15.1) // convert float64 in Gallons, carFuel is now type Gallons

	fmt.Println(busFuel, carFuel)

	//assigning carFuel to Gallons, raises Error
	// carFuel := Liters(23)
	fmt.Println(carFuel)

	// 1L = 0.264 Gallons
	// 1Gallon = 3.785L
	carFuel = Gallons(Liters(10.5) * 0.264)
	busFuel = Liters(Gallons(15.5) * 3.785)
	fmt.Printf("\nConverted \nCarFuel - %0.1f\nBusFuel - %0.1f\n", carFuel, busFuel)

	// defined type can be usedinoperations together with literal values
	fmt.Println("Operations with Literals - ")
	fmt.Println("carFuel - 2.3 = ", carFuel-2.3)
	fmt.Println("carFuel + 2.3 = ", carFuel+2.3)
	fmt.Println("carFuel * 2.3 = ", carFuel*2.3)

	// defined types cannot be used in operations together with values of a
	// different type, even if the other type has the same underlying type.
	// this is to protect developers from accidentally mixing the two types
	// fmt.Println("carFuel * busFuel", carFuel*busFuel)  // raises error
}
