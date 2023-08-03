package main

import (
	"fmt"
)

var x int

type person struct {
	first string
	last  string
}
type foo int

var y foo

const bar = 45

func main() {
	y = bar
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))

}
