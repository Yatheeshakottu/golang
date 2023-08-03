package main

import (
	"fmt"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			count++
		}
	}

	fmt.Println(count)
}
/*package main

import (
	"fmt"
)

func main() {
	var A, B string
	fmt.Scan(&A)
	fmt.Scan(&B)

	count := 0
	for i := 0; i < len(A); i++ {
		if A[i] == B[i] && A[i] == '1' {
			count++
		} else if A[i] == B[i] && A[i] == '0' {
			count++
		}
	}

	fmt.Println(count)
}
*/