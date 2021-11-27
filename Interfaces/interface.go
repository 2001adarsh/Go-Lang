package main

import "fmt"

//In Go interfaces, we define behaviour and in struct, we define data fields.

//method set with value types, must have all receivers as value types.
//but method set with pointer types, can have all receivers both value and pointer types.


func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("hello Go!"))

	cnt := IntCounter(0)
	var inc Incrementer = &cnt
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ,",inc.Increment())
	}
}

// Writer Naming convention: If interface has only 1 function, then it should have +er at last.
type Writer interface {
	Write([]byte) (int, error)
}
type ConsoleWriter struct{}

func (c ConsoleWriter) Write(bytes []byte) (int, error) {
	n, err := fmt.Println(string(bytes))
	return n, err
}

//Another Interface:
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (i *IntCounter) Increment() int {
	*i++
	return int(*i)
}
