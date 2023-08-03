package main

import (
	"fmt"
)

type i interface{}

func area(i interface{}) {
	fmt.Printf("Type  %T,value %v\n ", i, i)

}
func main() {
	s := "aytheesha"
	area(s)
	i := 07
	area(i)
}
