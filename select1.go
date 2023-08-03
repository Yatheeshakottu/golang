package main

import "fmt"

func main() {
	c := make(chan bool)
	go func() {
		for {
			select {
			case <-c:
				//return
				fmt.Println(c)
			default:

			}
		}
	}()
	c <- true

}
