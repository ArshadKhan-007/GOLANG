package main

import (
	"fmt"
	"maps"
)

func main() {
	//maps is a key value pair data structure
	// Creating a map

	m := make(map[string]int) //make function is used to create a map, here key is string and value is int

	//Adding an element to map
	m["one"] = 1
	m["two"] = 2

	//Accessing an element from map
	fmt.Println(m["one"]) //We have to use key to access value
	fmt.Println(m["two"])

	//Deleting an element from a map
	delete(m, "two")
	fmt.Println(m)

	//Getting length of a map
	fmt.Println(len(m))

	//Clear a whole map
	clear(m)
	fmt.Println(m)

	//Declaring a map without make function
	m1 := map[string]int{
		"three": 3,
		"four":  4,
	}

	// Check if a value exists in a map
	value, ok := m1["three"]

	if ok {
		fmt.Println("Does Exist")
	} else {
		fmt.Println("Doesn't exist")
	}
	fmt.Println(value)

	//Checking for equality
	m2 := map[string]int{"price": 30, "phone": 10}
	m3 := map[string]int{"price": 30, "phone": 10}

	fmt.Println(maps.Equal(m2, m3)) //Here, maps is

}
