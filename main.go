package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	i, j := 1, 5
	fmt.Println("i, j: ", i, j, "And address are at: ", &i, &j)
	p := &i
	fmt.Println(*p) // prints the value of the address of i
	p = &j
	fmt.Println(*p)        // prints the value of the address of j
	fmt.Println(square(5)) // prints sqaure of the number
	fmt.Println("Value of j before and after squareAdd", j, squareAdd(&j))
	var absoluteZero int = -459
	fmt.Println(absoluteZero)
	myBool := 5 > 8
	fmt.Println(myBool)

	/*
		Raw string literals are character sequences between back quotes, often called back ticks. Within the quotes,
		 any character will appear just as it is displayed between the back quotes, except for the back quote
		 character itself.Raw string literals may also be used to create multiline strings:
	*/
	a := `Say "hello" to Go!`
	fmt.Println(a)
	//arrays
	seaCreatures := []string{"shark", "cuttlefish", "squid", "mantis shrimp"}
	seaCreatures = append(seaCreatures, "adarsh")
	fmt.Println(seaCreatures)

	//maps
	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}
	fmt.Println(sammy)
}

func square(x int) int {
	return (x * x)
}

func squareAdd(p *int) int {
	*p *= *p
	return (*p)
}
