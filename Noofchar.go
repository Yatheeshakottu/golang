package main

import (
	"fmt"
)

func NoofChars(s string) {
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {

		mp[s[i]]++
	}

	for v, k := range mp {
		fmt.Printf("%c", v)
		fmt.Println("", k)
	}
	ml := 0
	var ch byte
	for v, k := range mp {
		if ml < k {
			ml = k
			ch = v
		}
	}
	fmt.Printf("max  reapeting letter %c", ch)
	fmt.Println("", ml)
}
func main() {
	NoofChars("Iam going to home right now")
}
