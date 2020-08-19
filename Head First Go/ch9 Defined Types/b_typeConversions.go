package main

import (
	"fmt"
)

type Liters float64  // defined type
type Gallons float64 // defined type

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func main() {
	carFuel := Gallons(1.2)
	// add 40L to carFuel
	carFuel += ToGallons(Liters(40))
	fmt.Println("carFuel After adding 40L - ", carFuel)

	busFuel := Liters(4.5)
	// add 40L to busFuel
	busFuel += ToLiters(Gallons(40))
	fmt.Println("busFuel After adding 40Gallons - ", busFuel)

}
