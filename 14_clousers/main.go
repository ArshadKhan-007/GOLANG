package main

import "fmt"

//Clousers are functions that capture the surrounding state or environment in which they were created
//In Go, closures are implemented using anonymous functions that can access variables from their enclosing scope
func counter() func() int { //here, counter returns a function that returns an integer
	var count int = 0
	return func() int { //Here, an anonymous function is returned that increments and returns the count variable
		count++
		return count
	}
}

func makeadder() func(int) int {
	var sum int = 0
	return func(x int) int {
		sum += x
		return sum

	}
}

func main() {
	increment := counter() //Here, increment is a closure that captures the count variable from the counter function
	fmt.Println(increment())
	fmt.Println(increment())

	adder := makeadder()
	fmt.Println(adder(5))
	fmt.Println(adder(10))
	fmt.Println(adder(20))
}
