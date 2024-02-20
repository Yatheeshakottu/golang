package main

import (
	"fmt"
)

func genperm(prefix string, remaining string) {
	if len(remaining) == 0 {
		fmt.Println(prefix)
	}
	for i := 0; i < len(remaining); i++ {
		newprefix := prefix + string(remaining[i])
		newremaining := remaining[:i] + remaining[i+1:]
		genperm(newprefix, newremaining)
	}

}
func main() {
	input := "jam"
	genperm("", input)
}
