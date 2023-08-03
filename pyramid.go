// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var k int
	//var rows = 5
	for i := 1; i <= 5; i++ {
		//k = 0
		for j := 1; j <= 5-i; j++ {
			fmt.Print(" ")
		}
		k = 0
		for {
			fmt.Print("*")
			k++
			if k == 2*i-1 {
				break
			}

			//fmt.Println()
		}
		fmt.Println("")
	}
}
