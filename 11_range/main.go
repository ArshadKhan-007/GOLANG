package main

import "fmt"

func main() {
	//range function is used to iterate over elements in a variety of data structures.
	//It can be used with arrays, slices, maps, strings, and channels.
	//When using range with arrays or slices, it returns both the index and the value of each element.
	//When used with maps, it returns the key and value pairs.
	//When used with strings, it returns the index and the corresponding rune (character).
	//When used with channels, it receives values from the channel until it is closed.

	// Example with slice
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers { //Here, range iterates over the slice 'numbers'. And it takes two variables: 'index' and 'value'.
		//In each iteration, 'index' holds the current index of the element, and 'value' holds the value of the element at that index.
		fmt.Println("Index:", index, "Value:", value)
	}

	// Example with map
	person := map[string]int{"Alice": 30, "Bob": 25, "Charlie": 35}
	for key, value := range person {
		fmt.Println("Key:", key, "Value:", value)
	}
	// Here, range iterates over the map 'person'. In each iteration, 'key' holds the current key (name), and 'value' holds the corresponding value (age).

	// Example with string
	str := "hello"
	for index, char := range str {
		fmt.Println("Index:", index, "Character:", char)
	}
	// Here, range iterates over the string 'str'. In each iteration, 'index' holds the current index of the character, and 'char' holds the corresponding rune (character).

	// Example with array
	arr := [3]int{10, 20, 30}
	for index, value := range arr {
		fmt.Println("Index:", index, "Value:", value)
	}
	// Here, range iterates over the array 'arr'. In each iteration, 'index' holds the current index of the element, and 'value' holds the value of the element at that index.
}
