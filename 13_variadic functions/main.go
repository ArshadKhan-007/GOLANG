package main

import "fmt"

func sum(nums ...int) int { //Here, nums is a slice of integers
	total := 0
	for _, num := range nums { //Iterating over the slice
		total += num
	}
	return total
}

// func add(nums ...interface{}) { //Here, interface{} is used to accept any type of arguments
// 	for _, num := range nums {
// 		fmt.Println(num)
// 	}
// }

func main() {
	// Variadic function example
	//Variadic functions is a function that can take a variable number of arguments which means you can pass any number of arguments to the function
	//In Go, you can define a variadic function by using ellipsis (...) before the type of the last parameter
	//In this example, the sum function takes a variable number of integers as arguments and returns their sum
	result := sum(2, 5, 8, 0) //Passing multiple integers
	fmt.Println(result)
	nums := []int{1, 2, 3, 4}
	fmt.Println(sum(nums...)) //Passing a slice of integers using ... operator

}
