package main

import (
	"fmt"
)

func main() {
	n := "yatheesha"
	switch n {
	case "yathee":
		fmt.Println("Print yathee")
	case "yatheesha":
		fmt.Println("print yatheesha")
	case "kottu":
		fmt.Println("Prints")
	default:
		fmt.Println("print default")
	}
}
