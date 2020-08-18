package main

import (
	"fmt"
)

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subscriber) {
	fmt.Println("Name - ", s.name)
	fmt.Println("Rate - ", s.rate)
	fmt.Println("Active? - ", s.active)
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func main() {
	sub1 := defaultSubscriber("Manav") // create defaultSubscriber
	printInfo(sub1)
	sub1.rate = 4.9 // update rate
	fmt.Printf("%#v\n", sub1)
	printInfo(sub1)
}
