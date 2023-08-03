package main

import (
	"fmt"
)

func permutations(str string) []string {
	var res []string
	if len(str) == 0 {
		res = append(res, "")
		return res
	}
	for i := 0; i < len(str); i++ {
		x := string(str[i])
		remainder := str[0:i] + str[i+1:]
		for _, v := range permutations(remainder) {
			res = append(res, x+v)
		}
	}
	return res

}
func main() {
	a := "ABC"
	fmt.Println(permutations(a))
}
