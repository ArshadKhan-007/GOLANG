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

func main() {
	result := add(2, 2)
	fmt.Println(result)

	product := subtract(10, 5)
	fmt.Println(product)

	fmt.Println(getLanguages())
}
