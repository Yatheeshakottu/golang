// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var large1, large2 int
	arr := []int{56, 34, 78, 987, 543}
	large1 = arr[0]
	for i := 0; i < len(arr); i++ {
		if large1 < arr[i] {
			large2 = large1
			large1 = arr[i]

		} else if large2 < arr[i] {
			large2 = large1
		}
	}
	fmt.Println(large2)
}
