package main

import (
	"fmt"
)

func main() {
	foo()
	bar("james")
	s1 := woo("missmoneypenny")
	fmt.Println(s1)
}
func foo() {
	fmt.Println("hello from foo")

}
func bar(s string) {
	fmt.Println("hello", s)
}
func woo(st string) string {

	return fmt.Sprint("hii from woo, ", st)

}
