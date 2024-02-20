package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wg.Add(1)
	go sum(n)
	//fmt.Println(n)
	wg.Wait()

}
func sum(n []int) {
	defer wg.Done()
	sums := 0

	for i := 0; i < len(n); i++ {
		square := n[i] * n[i]

		sums += square
	}

	fmt.Println(sums)
}

/*
package main

import "fmt"

func main() {
	sum := 0
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(n); i++ {
		square := n[i] * n[i]
		sum += square

	}
	fmt.Println(sum)
}




*/
