package main

import "fmt"

func checkanagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]]++
	}

	for j := 0; j < len(t); j++ {
		if _, ok := mp[t[j]]; !ok {
			return false
		}
		mp[t[j]]--
	}
	//for _, v := range mp {
	//if v != 0 {
	//	return false
	for i := 0; i < len(s); i++ {
		if mp[s[i]] != 0 {
			return false
		}
	}
	return true
}
func main() {
	input := checkanagram("rat", "bat")
	fmt.Println(input)
	output := checkanagram("listen", "silent")
	fmt.Println(output)

}
