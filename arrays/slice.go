package main

import "fmt"

func main() {
	//slices are similar to array, but they are naturally reference typed.
	//same as what array is in C++.
	//syntax: only need [] brackets to make them as slice without [...] (3 dots)
	a := []int{1, 2, 3}
	b := a
	b[1] = 5

	fmt.Println(a) //one of the slices changes the data, will have impact other places also.
	fmt.Println(b)

	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b = a[:]    //slice of all elements.
	c := a[3:]  //slice from 4th element to the end.
	d := a[:6]  //slice first 6 elements
	e := a[3:6] //slice the 4th, 5th and 6th element.
	fmt.Println(a)
	fmt.Println(b, c, d, e)
	//so basically it is [ )

	//another method to create slice: make function
	a = make([]int, 3, 100) // type, len, capacity
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("capacity: %v\n", cap(a))

	//push_back. -> increases capacity as same. --> may cause expensive copy operation.
	a = append(a, 12, 11, 423) //adding multiple elements at end.
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("capacity: %v\n", cap(a))

	//appending another slice in a.
	//a = append(a, []int{1,2,3,4})  -> just this is not going to work.
	a = append(a, []int{1, 2, 3, 4}...) //... are called spread operator, it spreads the element of slice into individual integers.

	//stack operations in slice.
	a = append(a) //push function.
	//but to pop elements.
	a = a[1:]        //this will pop the first element, basically recreate the array.
	a = a[:len(a)-1] //this will remove the element from end.
	fmt.Println(a)

	//suppose want to delete the third element.
	a = append(a[:2], a[3:]...)
	fmt.Println(a)

}
