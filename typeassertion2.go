package main

import (
	"fmt"
)

func main() {
	var value interface{} = 2001
	 var value1 int = value.(int)
	fmt.Println(value1)
	value2, test := value.(string)
	if test {
		fmt.Println("String value is found")
		fmt.Println(value2)
	} else {
		fmt.Println("String value is not found")
	}
}
