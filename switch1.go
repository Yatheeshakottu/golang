package main

import (
	"fmt"
)

func foo(i interface{}) {
	switch y := i.(type) {
	case int:
		fmt.Println("THe value is ", v)
	case string:
		fmt.Println("the string len:", v, len(v))
	default:
		fmt.Println("the floating value is", v)
	}

}
func main() {

	foo(35)
	foo("ytaheesha")
	foo(43.23)
}
