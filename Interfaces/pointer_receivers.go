package main

import "fmt"

//The syntax for defining methods on the pointer receiver is nearly identical to defining methods on the value receiver.
//The difference is prefixing the name of the type in the receiver declaration with an asterisk (*).

//The method sets for the pointer receiver and the value receiver are different because methods that receive a pointer
//can modify their receiver where those that receive a value cannot.

type Submersible interface {
	Dive()
}

type Shark struct {
	Name         string
	isUnderwater bool
}

func (s Shark) String() string {
	if s.isUnderwater {
		return fmt.Sprintf("%s is underwater", s.Name)
	}
	return fmt.Sprintf("%s is on the surface", s.Name)
}

func (s *Shark) Dive() {
	s.isUnderwater = true
}

func submerge(s Submersible) {
	s.Dive()
}

func main() {
	s := &Shark{
		Name: "Sammy",
	}

	fmt.Println(s)
	submerge(s)
	fmt.Println(s)
}
