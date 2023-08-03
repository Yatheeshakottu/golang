package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello playground")
	f := func() {
		fmt.Println("My first func expression")
	}
	f()
	g := func(x int) {
		fmt.Println("The big brother is watching movie:", x)
	}
	g(1987)
}
