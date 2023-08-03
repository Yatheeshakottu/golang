// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var k int
	var rows int = 5
	for i := 1; i <= rows; i++ {
		k = 0
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		//k = 0
		for {
			//k++
			fmt.Print("*")
			k++
			if k == 2*i-1 {
				break
			}
		}
		fmt.Println("")

	}
}
