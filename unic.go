package main

import (
	"fmt"
	"sync"
)

// func main(){
func main() {
	wg := sync.WaitGroup{}
	//fmt.Println("1")
	wg.Add(1)
	go sampleRoutine(&wg)
	wg.Wait()
	fmt.Println("1")
	//time.Sleep(1 * time.Second)
	fmt.Println("2")
}

func sampleRoutine(wg *sync.WaitGroup) {
	fmt.Println("3")
	wg.Done()
}
