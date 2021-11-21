package main

import (
	"fmt"
)

type Doctor struct { //if name starts with Uppercase -> exported to the package, else local to package.
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name   string `required_max:"100"` //tag
	Origin string `example:"origin"`
}
type Bird struct { //Inheritance is not present in GO, so Embedded property(Composition in JAVA) is used to replicate, but Bird is still not the type of Animal, it is still an independent struct.
	Animal
	SpeedKPH float32
	canFly   bool
}

func main() {
	aDoctor := Doctor{
		actorName:  "Adarsh Singh",
		number:     9,
		companions: []string{"Ashi", "Ayushi", "Anuj"},
	}
	fmt.Println(aDoctor, "Doctor Name:", aDoctor.actorName, "Brother:", aDoctor.companions[2])

	//anonymous structures, very rarely used.
	anonymous := struct {
		name string
		id   int
	}{name: "Adarsh", id: 1168}
	fmt.Println(anonymous)

	//Go doesn't has classes/inheritance, but has something similar called composition.
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.canFly = false
	b.SpeedKPH = 48
	fmt.Printf("%v, %T", b, b) //it is still the type Bird and not Animal, so no inheritance.

}
