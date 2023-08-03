// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func anagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
	}
	for j := 0; j < len(t); j++ {
		m[string(t[j])]--
	}
	for i := 0; i < len(s); i++ {
		if m[string(s[i])] != 0 {
			return false

		}
	}
	return true
}
func main() {
	x := anagram("got", "tog")
	y := anagram("got", "hog")
	fmt.Println(x)
	fmt.Println(y)

}
