package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := "3.14356"
	fmt.Println("before:", reflect.TypeOf(x))
	fmt.Println("the string value is:", x)
	a, _ := strconv.ParseFloat(x, 67)
	fmt.Println("after:", reflect.TypeOf(a))
	fmt.Println("the float value is:", a)
}
