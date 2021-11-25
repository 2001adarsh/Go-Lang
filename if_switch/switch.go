package main

import "fmt"

// In Go we have break in every case by default implicitly, no need to put it there,
// but if need is to run the statements below it, we have
// fallthrough.

func main() {
	flavors := []string{"chocolate", "vanilla", "strawberry", "banana"}

	for _, flav := range flavors {
		switch flav {
		case "strawberry":
			fmt.Println(flav, "is my favorite!")
			fallthrough //also includes the cases below it.
		case "vanilla", "chocolate":
			fmt.Println(flav, "is great!")
		default:
			fmt.Println("I've never tried", flav, "before")
		}
	}

	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i is dont know what type.")
	}

}
