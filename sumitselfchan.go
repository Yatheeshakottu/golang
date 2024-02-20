package main

import (
	"fmt"
	"sync"
)

func main() {
	input := []int{7, 4, 3, 10, 9}
	resultchan := make(chan int)
	var wg sync.WaitGroup
	for _, num := range input {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			resultchan <- n + n
		}(num)

	}
	go func() {
		wg.Wait()
		close(resultchan)

	}()
	results := make([]int, len(input))
	i := 0
	for num := range resultchan {
		results[i] = num
		i++

	}
	fmt.Println(results)
}
