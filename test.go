// add_test.go

package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}

//func main() {
//TestAdd(&result)
//}
