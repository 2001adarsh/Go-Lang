package main

import "fmt"

//Enumerated constants.
const (
	a = iota //iota means counter.
	b
	c
)

const (
	_ = iota + 5 //right only variable
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	_ = iota //ignore first value by assigning blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

func main(){
	const myConst = 42
	const str string = "Hello"
	const fl float32 = 1.72
	fmt.Printf("%v %T %T %T\n", myConst, myConst, str, fl)

	fmt.Println(a, b, c, catSpecialist)

	fileSize := 40000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
}
