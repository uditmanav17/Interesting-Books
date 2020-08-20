// previously, we could have accessed date's attributes using date.Year and completely bypassing setters.
// Now, by moving date to different package and making its attributes unexported,
// we can only acces them using setters
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
}
