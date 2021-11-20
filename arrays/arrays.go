package main

import "fmt"

func main() {
	grades := [3]int {97, 98, 99}//creating a size of 3 array
	fmt.Printf("Grades: %v, %T\n", grades, grades)

	hold := [...]int {78, 23, 43} // ... means that create a size of array just enough to hold all elements.
	fmt.Printf("Hold : %v %v\n", hold, len(hold))

	var arr [3]string //creating an array.
	arr[0] = "Adarsh"
	arr[1] = "Singh"
	arr[2] = "Great"
	fmt.Println(arr)
	fmt.Println(arr[1])

	//multidimensional array:
	var matrix [3][3]int = [3][3]int{ [3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1} }
	fmt.Println(matrix)

	copy := arr // In GO, copy will not be pointing to elements of arr but rather will be creating same array again (copy by value)
	arr[2] = "change wont be reflected"
	fmt.Println(copy)

	//need to use pointer for *copy by reference
	copyRef := &arr
	arr[2] = "change to be refelected"
	fmt.Println(*copyRef)
}