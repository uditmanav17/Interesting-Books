package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/calendar"
)

func main() {
	event := calendar.Event2{}
	event.SetTitle("Birthday")

	err := event.SetYear(2020)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetMonth(11)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Event Title -", event.Title())
	fmt.Println("Event Day -", event.Day())
	fmt.Println("Event Month -", event.Month())
	fmt.Println("Event Year -", event.Year())

	fmt.Println(event)
}
