package main

import "fmt"

type motorvehicle interface {
	milage() float64
}
type BMW struct {
	distance     float64
	fuel         float64
	averagespeed string
}

func (b BMW) AverageSpeed() string {
	return b.averagespeed
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
func totalmilage(m motorvehicle) {
	au := m.(BMW)
	fmt.Printf(au.AverageSpeed())
}
func main() {
	b1 := BMW{
		distance:     976.5,
		fuel:         768.6,
		averagespeed: "76",
	}
	totalmilage(b1)
}
