package main

import "fmt"

// Generics in Go allow you to write functions and data structures that can operate on different types without sacrificing type safety.
// In simple terms, generics enable you to create reusable components that can work with any data type.

// Generic Function Example
// This function can take a slice of any type and print its elements
func printSlice[T any](items []T) { // T is a type parameter that can be any type
	for _, item := range items {
		fmt.Println(item)
	}
}

// Function specifically for int type and string type
// This function will only take slices of int or string types
func printAnotherSlice[T int | string](items []T) { // T is constrained to int or string types
	for _, item := range items {
		fmt.Println(item)
	}
}

// Struct with Generics Example
type Pair[T any] struct {
	elements []T
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}                         // Slice of integers
	stringSlice := []string{"Hello", "Generics", "in", "Go"} // Slice of strings

	// Calling the generic function with different types
	printSlice(intSlice)
	fmt.Println("-----")
	printSlice(stringSlice)
	fmt.Println("-----")
	printAnotherSlice(intSlice)
	fmt.Println("-----")
	printAnotherSlice(stringSlice)
	fmt.Println("----------")

	// Using the generic struct with different types
	p := Pair[int]{
		elements: []int{1, 24, 5},
	}
	fmt.Println(p)
	x := Pair[string]{
		elements: []string{"a", "b", "c"},
	}
	fmt.Println(x)
}

// As we can see in the above code example, the function printSlice is defined with a type parameter T.
// We haven't created new function for int and string types; instead, we use the same function for both types.
// This demonstrates the power of generics in Go, allowing for code reuse and type safety.
