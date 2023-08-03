package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%b ", i%2)
			//fmt.Printf("%d ", j%2)
		}

		fmt.Println(" ")
	}
}

/*package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 5; i++ {
		for j := i; j <= 5; j++ {
			fmt.Print(i%2, " ")
		}
		fmt.Println()

	}

}*/
