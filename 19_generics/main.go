package main

import "fmt"

// Generics in Go allow you to write functions and data structures that can operate on different types without sacrificing type safety.
// In simple terms, generics enable you to create reusable components that can work with any data type.

// Generic Function Example
func printSlice[T any](items []T) { // T is a type parameter that can be any type
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}                         // Slice of integers
	stringSlice := []string{"Hello", "Generics", "in", "Go"} // Slice of strings

	// Calling the generic function with different types
	printSlice(intSlice)
	fmt.Println("-----")
	printSlice(stringSlice)
}

// As we can see in the above code example, the function printSlice is defined with a type parameter T.
// We haven't created new function for int and string types; instead, we use the same function for both types.
// This demonstrates the power of generics in Go, allowing for code reuse and type safety.
