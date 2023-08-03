package main

import (
	"fmt"
)

type motorvehicle interface {
	milage() float64
}
type BMW struct {
	distance     float64
	fuel         float64
	averagespeed string
}
type Audi struct {
	distance float64
	fuel     float64
}

func (b BMW) milage() float64 {
	return b.distance / b.fuel
}
func (a Audi) milage() float64 {
	return a.distance / a.fuel
}

func totalmilage(m []motorvehicle) {
	tm := 0.0
	for _, v := range m {
		tm = tm + v.milage()
	}
	fmt.Printf("the total milage per month %f", tm)
}
func totalmilage(m)
func main() {
	s := BMW{
		distance:     899,
		fuel:         653,
		averagespeed: "36",
	}
	a1 := Audi{
		distance: 658,
		fuel:     763,
	}
	person := []motorvehicle{s, a1}
	totalmilage(person)
}
