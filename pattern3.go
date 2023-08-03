package main

import (
	"fmt"
)

func main() {
	for i := 5; i > 0; i-- {
		for j := 5; j > 0; j-- {
			//fmt.Printf("%d", j)
			fmt.Printf("%d", i)
		}
		fmt.Println()
	}
}

/*// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	for i := 5; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}*/
