package main

import (
	"fmt"
)

func reverse(str string) string {
	rn := []rune(str) //converting  in to rune
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		rn[i], rn[j] = rn[j], rn[i]
	}
	return string(rn)
}

func main() {
	str := "sujatha"
	strRev := reverse(str)
	fmt.Println(str)
	fmt.Println(strRev)
}
