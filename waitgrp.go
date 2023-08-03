package main

import (
	"fmt"
	"sync"
	"time"
)

// func main(){
func main() {
	wg := sync.WaitGroup{}
	fmt.Println("1")
	//	wg.Add(1)
	go sampleRoutine(&wg)
	//wg.Wait()

	time.Sleep(1 * time.Second)
	fmt.Println("2")
}

func sampleRoutine(wg *sync.WaitGroup) {
	//	defer wg.Done()
	fmt.Println("3")

}
