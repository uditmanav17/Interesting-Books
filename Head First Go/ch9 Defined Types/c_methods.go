package main

import (
	"fmt"
)

type Liters float64  // defined type
type Gallons float64 // defined type
type MilliLtrs float64

// this is a method
// params before func name
func (l Liters) ToGallons() Gallons { //method for ToGallons Liters type
	return Gallons(l * 0.264)
}

func (ml MilliLtrs) ToGallons() Gallons { //method for ToGallons MilliLtrs type
	return Liters(ml * 1000).ToGallons()
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func main() {
	carFuel := Gallons(1.2)
	fmt.Println("carFuel in Ltrs - ", carFuel.ToLiters())

	busFuel := Liters(4.5)
	fmt.Println("busFuel in Gallons - ", busFuel.ToGallons())

	mls := MilliLtrs(5000)
	Ltrs := Liters(5)
	fmt.Printf("%0.2f mls 2 Gallons is %0.2f\n", mls, mls.ToGallons())
	fmt.Printf("%0.2f Ltrs 2 Gallons is %0.2f\n", Ltrs, mls.ToGallons())
	fmt.Println("Is mls == Ltrs?", mls.ToGallons() == Ltrs.ToGallons())

}
