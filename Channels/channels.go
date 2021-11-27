package main

import (
	"fmt"
	"sync"
	"time"
)

//You can send values into channels from one goroutine and receive those values into another goroutine.
//Channels are the pipes that connect concurrent goroutines.

var wg = sync.WaitGroup{}

func main() {
	//make(chan int, 50) -> will create a buffer, so that channel can store upto 50 integers. by default there is only 1.

	ch := make(chan int) // chan to create channel with data type that it will pass. Hence, it is strongly typed.
	wg.Add(2)
	go func() {
		i := <-ch // will be receiving the value from channel.
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		ch <- 42 // getting data in the channel
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("Another example:")
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go func() {
			data := <-ch // will be receiving the value from channel.
			fmt.Println(data)
			wg.Done()
		}()
		go func(i int) {
			ch <- i // sending data in the channel -> this particular line will pause the execution until there is space in channel. Hence Deadlock might happen.
			wg.Done()
		}(i)
		wg.Wait()
	}

	//SELECT STATEMENTS
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
LOOP_OUT:
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
			break LOOP_OUT// to get out of the loop.
		}
	}

}

//In func definition we can define if it is receive only channel or read only channel by
// func( ch <-chan int) -> read only channel
// func( ch chan<- int) -> receive only channel

//select to select the channel and generally getting out of for loop of channel.
