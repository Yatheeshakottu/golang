package main

import (
	"fmt"
)

func slice(element []string) {
	element[1] = "yatheesha"
	fmt.Println("The modified output is:", element)
}
func main() {
	s := []string{"c", "Go", "java", "python"}
	fmt.Println("the intial output is:", s)
	slice(s)
	fmt.Println("The final output is:", s)
}
