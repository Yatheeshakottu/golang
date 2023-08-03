// given 2 arrays arr1[] nd arr2[] ,for each element in arr1[] find the same element in arr2[] and print the index value ,if not found print "NA”.
// Assume arr1[] and arrr2[] doesn't have duplicate values in Golang
package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{5, 4, 3, 7, 2, 1}

	for i := 0; i < len(arr1); i++ {
		found := false
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				fmt.Println(arr1[i], ":", j)
				found = true
				break
			}
		}
		if !found {
			fmt.Println(arr1[i], ": NA")
		}
	}
}
