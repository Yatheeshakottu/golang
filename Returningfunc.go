package main

import (
	"fmt"
)

func main() {
	s1 := foo()
	fmt.Println(s1)

}
func foo() string {

	s := "hello world"
	return s
}
