package main

import (
	"fmt"
)

func main() {
	var value interface{} = "Golang"
	value1 := value.(string)
	fmt.Println(value1)
	value2 := value.(int)
	fmt.Println(value2)
}
