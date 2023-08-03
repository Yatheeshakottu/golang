package main

import (
	"fmt"
)

func checkanagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {

		mp[s[i]]++
	}
	//fmt.Println(mp)
	for j := 0; j < len(s); j++ {
		if _, ok := mp[s[j]]; !ok {
			return false
		}
		mp[t[j]]--
	}

	for _, v := range mp {
		if v != 0 {
			return false
		}
	}
	return true
}
func main() {
	x := checkanagram("bat", "cat")
	t := checkanagram("listen", "silent")
	fmt.Println(x)
	fmt.Println(t)
}
