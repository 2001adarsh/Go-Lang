package main

import "fmt"

func main()  {
	var n bool = false
	fmt.Printf("%v , %T \n", n, n)

	//int -> uint8, uint16, uint32

	a := 10 // 1010
	b := 3 //  0011
	fmt.Println(a &^ b) // 0100 - both bit should be unset.
	fmt.Println( a << 3)

	//complex numbers:
	var comp complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", comp, comp)
	fmt.Printf("%v, %T\n", real(comp), real(comp))
	fmt.Printf("%v, %T\n", imag(comp), imag(comp))

	//Strings;
	//1. immutable, 2. s[2] will give byte number. 3.string concat 4.
	s := "this is a string"
	by := []byte(s) //slice of bytes - a lot of functions in go uses byte slices.
	fmt.Printf("%v , %T\n", by, by)

	//r := 'a'
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}