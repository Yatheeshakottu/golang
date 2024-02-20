package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GOroutines\t\t", runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go bar()
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GOroutines\t\t", runtime.NumGoroutine())
	wg.Wait()

}
func foo() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("even numbers:", i)
		}
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
	wg.Done()
}

/*A WaitGroup is designed to wait for a specific number of "tasks" to complete before allowing the program to continue. In the context of goroutines, a "task" is typically defined as a single unit of work that you want to wait for. In your code, each iteration of the loops in both foo() and bar() represents a unit of work. As such, each iteration should be treated as a separate task when using a WaitGroup.

However, in your code, you've placed the wg.Done() call inside the loop in the foo() function. This means that you are calling wg.Done() for each iteration of the loop, which causes the WaitGroup counter to decrease multiple times in rapid succession. This is not how a WaitGroup is intended to be used.

To correctly use a WaitGroup in this scenario, you should call wg.Done() once per task, not per iteration of the loop. Here's the corrected version of your code:
*/
//This is for when then two functions work is same then that no need to add extra counter value
/*func foo(){
for i:=0;i<10;i++{
fmt.Println("foo",i)
}
wg.Done()
}
func bar(){
	for i:=0;i<10;i++{
fmt.Println("bar",i)
	}
}*/
