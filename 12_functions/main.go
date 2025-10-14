package main

import "fmt"

//functions
func add(x int, y int) int {
	return x + y
}

func subtract(a, b int) int { // shorthand for same type
	return a - b
}

func getLanguages() (string, string) { // multiple return values
	return "Go", "JavaScript"
}

//function as an argument to another function
func processIt(fn func(a int) int) { //Here, processIt is a function that takes another function fn as an argument. The fn function takes an int as input and returns an int.
	// Inside processIt, we call the fn function with the argument 5 and ignore the return value.
	fn(5)
}

func main() {
	result := add(2, 2)
	fmt.Println(result)

	product := subtract(10, 5)
	fmt.Println(product)

	fmt.Println(getLanguages())

	//anonymous function
	fn := func(a int) int { //Here, we define an anonymous function and assign it to the variable fn.
		return 2
	}
	processIt(fn)
}
