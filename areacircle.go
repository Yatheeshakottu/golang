package main

import (
	"fmt"
	"math"
)

func main() {
	var area, radius float64

	//length := 8
	radius = 10
	area = math.Pi * (radius * radius)
	fmt.Println(area)
}
