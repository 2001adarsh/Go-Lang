package main

import (
	"fmt"
)

//errors, which are unanticipated by the programmer

//Go doesn't have Exceptions, but most of the functions return err, if something happens. So no need of exceptions.
//But there can be cases like divide by 0, where things can change --> in those situations we can use PANIC.

func main() {
	//a, b := 1, 0
	//ans := a/b // divide by 0 panic
	//fmt.Println(ans)

	//s := &Shark{"Sammy"}
	//s = nil
	//s.SayHello() //nil pointer panic

	//panic("oh no!") //panic function

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("Hello Go"))
	//})
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	panic(err.Error())
	//}

	fmt.Println("start")
	defer fmt.Println("This was deferred") //Defer will run before panic, to release the holdings. It will execute before panic.
	panic("something bad happened")
	fmt.Println("end")

}

type Shark struct {
	Name string
}

func (s *Shark) SayHello() {
	fmt.Println("Hi! My name is", s.Name)
}
