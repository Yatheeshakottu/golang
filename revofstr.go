package main


import (
	"fmt"
)

func reverse(s string) string {
	rn := []rune(s)
	for i := 0; i < len(rn); i++ {
		for j := 0; j < len(rn); j++ {
			//if rn[i] > rn[j] {
			if rn[i] < rn[j] {
				rn[i], rn[j] = rn[j], rn[i]
			}
		}
	}
	return string(rn)
}
func main() {
	s := "yatheesha"
	srev := reverse(s)
	fmt.Println(srev)

}
