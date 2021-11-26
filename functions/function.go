package main

import "fmt"

func main() {
	//variadic functions.
	result := sum("The values are:", 1, 2, 3, 4, 6)
	fmt.Println("Sum: ", *result)

	//named return func
	m := multiply(5, 6)
	fmt.Println(m)

	//return err from function.
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)

	//Anonymous function.
	f := func() {
		fmt.Println("Hello GO!")
	} //() //these are to invoke the func
	f()

	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i) //This will run async, and not passing the value, func will take value from outer scope, and since it is async, it's result may vary.
	}



}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func multiply(a, b int) (result int) { //(result int) -> named return, this variable will already be present beneath
	result = a * b
	return
}

func sum(msg string, values ...int) *int {
	fmt.Printf("%v, %T\n", values, values)
	result := 0
	for _, val := range values {
		result += val
	}
	return &result //Ability to return address of variable for performance benefit.
}
