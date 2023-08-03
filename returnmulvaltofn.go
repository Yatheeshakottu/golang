package main

import (
	"fmt"
)

func reverse(a, b, c, d string) (string, string, string, string) {
	return a, c, b, d
}
func main() {
	x, y, z, a := reverse("WELCOME", "TO", "THE", "GOPLAYGROUND")
	fmt.Println(x, y, z, a)
}
