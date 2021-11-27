package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{} //as many can read, but only one can write at a time.

func main() {
	fmt.Println("Number of threads present: ", runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(100) // increasing something like this will cause other problems, increasing memory overheads and scheduler has to work harder and perfomance will fall.

	fmt.Println("Number of threads present: ", runtime.GOMAXPROCS(-1))

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock() // read lock
		go sayHello()
		m.Lock() // write lock
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Println("Hello ", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
