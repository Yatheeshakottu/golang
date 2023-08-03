package main

import (
	"fmt"
)

func main() {

	n, err := fmt.Println("yatheesha")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
