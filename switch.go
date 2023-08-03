package main

import (
	"fmt"
)
//switch statement is similar to if else ,but it will stop once the condition is satisfied.it is not carry forward to the next statement after valid pass.
func main() {
	n := true
	switch n {
	case false:
		fmt.Println("print its false")
		//fallthrough
	case (7 == 7):
		fmt.Println("print its true")
		fallthrough
	case (7 != 8):
		fmt.Println("print the value")
	//	fallthrough
	case (7 == 7):
		fmt.Println("print its a true")

	}
}
