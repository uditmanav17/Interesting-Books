package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/calendar"
)

func main() {
	var newDate calendar.Date

	err := newDate.SetYear(2020)
	if err != nil {
		log.Fatal(err)
	}

	err = newDate.SetMonth(11)
	if err != nil {
		log.Fatal(err)
	}

	err = newDate.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newDate)

	// get date using getters
	fmt.Println("Day -", newDate.Day())
	fmt.Println("Month -", newDate.Month())
	fmt.Println("Year -", newDate.Year())
}
