package main

import "fmt"

func main() {
	var a, b string
	fmt.Print("enter the strings:")
	fmt.Scan(&a, &b)
	op1 := ""
	op2 := ""
	for i := 0; i < len(a); i++ {
		if !contains(b, a[i]) {
			op1 += string(a[i])
		}
	}

	for j := 0; j < len(b); j++ {
		if !contains(a, b[j]) {
			op2 += string(b[j])
		}
	}
	fmt.Println("op1", op1)
	fmt.Println("op2", op2)
}

func contains(a string, b byte) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == b {
			return true
		}
	}
	return false

}
