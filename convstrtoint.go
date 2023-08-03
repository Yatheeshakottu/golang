package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := "yatheesha"
	fmt.Println("Before:", reflect.TypeOf(x))
	fmt.Println("the string value is:", x)
	a, _ := strconv.ParseInt(x, 1, 65) //we can give any number of integer values
	fmt.Println("after:", reflect.TypeOf(a))
	fmt.Println("The interger value is:", a)
}
