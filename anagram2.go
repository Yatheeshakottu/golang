package main

import (
	"fmt"
)

func checkanagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp := make(map[string]int)
	for i := 0; i < len(s); i++ {
		mp[string(s[i])]++
	}
	for j := 0; j < len(t); j++ {
		mp[string(t[j])]--
	}
	for i := 0; i < len(s); i++ {
		if mp[string(s[i])] != 0 {
			return false
		}
	}
	return true
}
func main() {
	x := checkanagram("yathee", "ethaet")
	t := checkanagram("listen", "silent")
	fmt.Println(x)
	fmt.Println(t)
}
