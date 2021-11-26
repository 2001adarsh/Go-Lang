package main

import "fmt"

func main() {
	s := []int{11, 22, 33}

	// k -> index, v -> value. - for non-map collections.
	for k, v := range s {
		fmt.Println(k, v)
	}

	//k -> key, v -> value -- for map collection.
	statePopulation := map[string]int{
		"UP":        324524,
		"Rajasthan": 22345,
		"Kerela":    1234,
		"Delhi":     23453,
	}
	for k, v := range statePopulation{
		fmt.Println(k, v)
	}

	//if only need values.
	for _, v := range s {
		fmt.Printf("%v ", v)
	}

}
