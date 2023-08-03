package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		m.Lock()
		go func(i int) {
			//defer wg.Done()
			m.Unlock()
			fmt.Println(i+1, i)
			wg.Done()
		}(i)

	}
	wg.Wait()
	//time.Sleep(10000)
}
