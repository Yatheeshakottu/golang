package main

import "fmt"

func main() {
	s := "yatheesha"
	rn := []rune(s)
	for i, j := 0, len(rn)-1; i < j; i, j = i+1, j-1 {
		rn[i], rn[j] = rn[j], rn[i]
	}
	fmt.Println(string(rn))
}
