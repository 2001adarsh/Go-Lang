package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil { // after panic, if recover() func gives err, then we can actually do something and stop this function execution.
			log.Println("Error: ", err) //Print and close the function, will not stop the application.
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking") // If panic happened, then this function will not work any further.
}
