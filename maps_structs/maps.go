package main

import "fmt"

func main() {
	//basic initialization
	statePopulation := map[string]int{
		"UP":        324524,
		"Rajasthan": 22345,
		"Kerela":    1234,
		"Delhi":     23453,
	}
	fmt.Println(statePopulation, statePopulation["Delhi"], statePopulation["Adarsh"]) //by default the value is 0 if not present.

	//maps of slices cannot be made but can be of an array
	m := make(map[[3]int]int) //initialization using make function.
	m[[3]int{1, 2, 3}] = 4    //updating values.
	fmt.Println(m)

	c := make(map[string]string)
	fmt.Println(c["Adarsh"]) //doesn't print anything.
	fmt.Println("A line before gone")
	//this brings inconsistency
	pop, ok := c["Adarsh"]
	fmt.Println("Ok will tell if map has that value or not?", ok, pop)

	//return order of map is not guaranteed.
	//-> Similar to HashMap of JAVA

	//delete from map.
	delete(statePopulation, "Delhi")
	fmt.Println(statePopulation)

	for key, value := range statePopulation {
		fmt.Printf("%q is the key for the value %q\n", key, value)
	}

	//Existence checking
	if count, ok := statePopulation["sammy"]; ok {
		fmt.Printf("Sammy has a count of %d\n", count)
	} else {
		fmt.Println("Sammy was not found")
	}

	//clear
	//declare variable of same map type and garbage collector will clear the map.
	statePopulation = map[string]int{}
	fmt.Println("Map cleared: ", statePopulation)
}
