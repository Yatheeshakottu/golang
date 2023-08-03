package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Parsebool is used to convert a string type in to a bool value
func main() {
	s := "yatheesha"
	fmt.Println("before:", reflect.TypeOf(s))
	fmt.Println("string value is:", s)
	b, _ := strconv.ParseBool(s)
	fmt.Println("after:", reflect.TypeOf(b))
	fmt.Println("boolean  value is:", b)
	x := "t"
	fmt.Println("before:", reflect.TypeOf(x))
	fmt.Println("string value is:", x)
	a, _ := strconv.ParseBool(x)
	fmt.Println("after:", reflect.TypeOf(a))
	fmt.Println("the bool value is:", a)
}
