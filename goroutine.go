//package external
//func myfunction(){

//}
package main

import (
	"fmt"
	"sync"
	"external"
)

var (
	counter = 0
	mutex   sync.Mutex
)
var wg sync.WaitGroup

func main() {
	//external.myfunction()
	//wg.Add(1)
	//sum(wg, &i)
	//wg.Wait()
	//func sum(wg *sync.WaitGroup) {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrementcounter()
		}()
	}
	wg.Wait()
	fmt.Println("counter value", counter)
}
func incrementcounter() {
	mutex.Lock()
	defer mutex.Unlock()
	counter++
}
