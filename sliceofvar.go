package main

import (
	"fmt"
)

func main() {
	arr := []string{"I", "am", "Going", "to", "office"}
	fmt.Println(arr)
	fmt.Println(cap(arr))
	fmt.Println(len(arr))
	slicearr := arr[0:3]
	fmt.Println(slicearr)
	fmt.Println(len(slicearr))
	fmt.Println(cap(slicearr))

}
