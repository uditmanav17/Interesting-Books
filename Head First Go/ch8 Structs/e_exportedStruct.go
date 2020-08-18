package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func printInfo(s *magazine.Subscriber) {
	fmt.Println("Name - ", s.Name)
	fmt.Println("Rate - ", s.Rate)
	fmt.Println("Active? - ", s.Active)
}

func defaultSubscriber(name string) *magazine.Subscriber {
	var s magazine.Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}

func main() {
	sub1 := defaultSubscriber("Manav") // create defaultSubscriber
	printInfo(sub1)
	sub1.Rate = 4.9 // update rate
	fmt.Printf("\n%#v\n\n", sub1)
	printInfo(sub1)

	sub2 := magazine.Subscriber{Name: "Udit", Rate: 9.9, Active: true}
	printInfo(&sub2)
}
