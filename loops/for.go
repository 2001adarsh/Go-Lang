package main

import "fmt"

func main() {
	//simple way.
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for i:=0, j:=0; i<5; i++, j++  --> this is wrong in GO, as GO doesn't have ',' operator. to acheive this we can
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	//No While loop, instead can do this.
	i := 0
	for i < 5 {
		fmt.Printf("%v ", i)
		i++
	}

	// for DO WHILE loop, GO has
	i = 0
	for {
		fmt.Printf("%v ", i)
		i++
		if i == 5 {
			break
		}
	}

	// continue keyword is also present.
	//Breaking out of complete loop
FOR_LOOP_VAR:
	for p := 1; p <= 3; p++ {
		for q := 1; q <= 3; q++ {
			fmt.Println(p * q)
			if p*q >= 3 {
				break FOR_LOOP_VAR //will break out whole for loop.
			}
		}
	}


}
