// There is no guarantee on which goroutine will execute first. It is determined by the Go scheduler, which schedules goroutines based on their priority, and also takes into account other factors like thread availability, resource usage, etc. Therefore, it is unpredictable which function will start first in this program.
package main

import (
	"fmt"
	"sync"
)

func abc(s string, ch chan int, wg *sync.WaitGroup) {
	fmt.Println(s)
	wg.Done()
	ch <- 21
}
func xyz(a string, ch chan int, wg *sync.WaitGroup) {
	fmt.Println(a)
	wg.Done()
	ch <- 43
}
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go xyz("the first loop", ch, &wg)
	fmt.Println(<-ch)
	go abc("the second loop", ch, &wg)
	//wg.Wait()
	//time.Sleep(1)
	fmt.Println(<-ch)
	wg.Wait()
}
