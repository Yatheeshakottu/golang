package main

import (
	"bytes"
	"fmt"
)

func main() {
	str1 := []byte{'y', 'a', 't', 'h', 'e', 'e'}
	str2 := []byte{'s', 'h', 'a'}
	//bytes of num is lowercase values
	fmt.Println(str1, str2)
	x1 := bytes.ToUpper(str1)
	x2 := bytes.ToUpper(str2)
	//Toupper means uppercase values
	fmt.Println(x1)
	fmt.Println(x2)
}
