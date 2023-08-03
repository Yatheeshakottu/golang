package main

import (
	"fmt"
	"regexp"
)

func main() {

	str := "Golang regular expressions example"

	match, err := regexp.MatchString(`^G`, str)
	fmt.Println("Match: ", match, " Error: ", err)

}
