package main

import "time"

// Structs in Go are used to group related data together to form a single entity. They are similar to classes in other programming languages but do not have methods associated with them.

// A struct is defined using the `type` keyword followed by the name of the struct and the `struct` keyword. The fields of the struct are defined within curly braces.

type order struct {
	orderId   int
	amount    float64
	status    string
	createdAt time.Time
}

func main() {

}
