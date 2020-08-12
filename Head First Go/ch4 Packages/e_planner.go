package main

import (
	"dates"
	"fmt"
)

func main() {
	days := 3
	fmt.Println("Appointment in", days, "days")
	fmt.Println("Appointment in", days+dates.DaysInWeek, "days")
	// dates.DaysInWeek = 5

}
