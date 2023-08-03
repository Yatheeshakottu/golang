package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64 = 2, 4, 2
	var discriminant, root1, root2, imaginary float64
	discriminant = (b * b) - (4 * a * c)

	if discriminant > 0 {
		root1 = (-b + math.Sqrt(discriminant)/(2*a))
		root2 = (-b - math.Sqrt(discriminant)/(2*a))
		fmt.Println("the two roots are distinct", root1, root2)
	} else if discriminant == 0 {
		root1 = -b / (2 * a)
		root2 = -b / (2 * a)
		fmt.Println("The roots real and equal:", root1, root2)
	} else if discriminant < 0 {
		root1 = -b / (2 * a)
		root2 = -b / (2 * a)
		imaginary = math.Sqrt(-discriminant) / (2 * a)
		fmt.Println("the roots are imaginary ", root1, "+", " imaginary", root2, "+", "imaginary", imaginary)
	}
}
