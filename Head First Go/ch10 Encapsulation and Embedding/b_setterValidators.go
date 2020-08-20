package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("Invalid Year!")
	}
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid Month!")
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid Day!")
	}
	d.Day = day
	return nil
}

func main() {
	var newDate Date
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
