package main

import "fmt"

//whose purpose is to operate on instances of some specific type, called a receiver
//the addition of an extra parameter after the func keyword for specifying the receiver of the method.
//The receiver is a declaration of the type that you wish to define the method on.

//In short -> more like a func of a particular class in our case Struct.

type greeter struct {
	greeting string
	name     string
}

func (receiver greeter) greet() {
	fmt.Println(receiver.name, receiver.greeting)
}

func main() {
	g := greeter{
		greeting: "Hey how ya doing?",
		name:     "Adarsh",
	}
	g.greet()

}

// the receiver of method invocations is typically referred to by a keyword (e.g. this or self).
