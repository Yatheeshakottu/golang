package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type circle struct {
	radius float64
}
type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (s square) area() float64 {
	return s.length * s.length
}
func intro(s shape) {
	x := s.area()
	fmt.Println(x)
}
func main() {
	circl := circle{
		radius: 12.345,
	}
	squa := square{
		length: 15,
	}
	intro(circl)
	intro(squa)
	//fmt.Println(area)

}
