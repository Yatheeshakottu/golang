package main

import (
	"fmt"
)

func pointer(ival int) {
	ival = 0
}
func value(ival *int) {
	*ival = 0

}
func main() {
	i := 1
	fmt.Println(i)
	pointer(i)
	fmt.Println(i)
	value(&i)
	fmt.Println(i)
	value(&i)
	fmt.Println("value;", &i)
}
