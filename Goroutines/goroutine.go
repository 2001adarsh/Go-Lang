package main

import (
	"fmt"
)

//Do not communicate by sharing memory; instead, share memory by communicating.

//A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space. It is lightweight, costing little more than the allocation of stack space. And the stacks start small, so they are cheap, and grow by allocating (and freeing) heap storage as required.

//WaitGroup: To wait for multiple goroutines to finish, we can use a wait group.

//var wg = sync.WaitGroup{}
//var counter = 0

func main() {
	//Example 1:
	msg := "Hello"
	//wg.Add(1)
	go func() {
		//defer wg.Done()
		fmt.Println(msg)
	}()
	msg = "GO"
	//time.Sleep(100* time.Millisecond)
	//wg.Wait()

	//Example 2:
	//for i := 0; i < 10; i++ {
	//	wg.Add(2)
	//	go sayHello()
	//	go increment()
	//}
	//wg.Wait()
}
//
//func sayHello() {
//	fmt.Printf("Hello #%v\n", counter)
//	wg.Done()
//}
//func increment() {
//	counter++
//	wg.Done()
//}
