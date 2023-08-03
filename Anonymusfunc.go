package main

import (
	"fmt"
)

func main() {

	func() {
		fmt.Println("Anonymus func ran")
	}()
	func(x string) {
		fmt.Println("The meaning of life:", x)
	}("yatheesha")

}
