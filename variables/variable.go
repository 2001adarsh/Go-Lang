package main

import (
	"fmt"
	"strconv"
)

var packageLevel int = 12
//lower case variables are scoped to package.

var (
	more int = 10
	variables string = "hello"
	inside float32 = 45.34
	block bool = false
)

var UPPERCASEVAR int = 45 //exported from package, globally visible.
//written as java conventions.

func main(){
	// variable here takes precedence, (inner most packaging - shadowing).
	var i int = 42 //useful when declare a variable, but not initialize it.   //block scope.
	j := 52                   //to let compiler decide the type of variable.
	var theURL string = "http://google.com"
	fmt.Printf("%v %T\n", packageLevel, j)
	fmt.Println(i, j, theURL)

	//converting the variables.
	k := float32(i)
	fmt.Printf("%v %T\n", k, k)

	str := strconv.Itoa(i)
	fmt.Println(str)

}

