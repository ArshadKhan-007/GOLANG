package main

import "fmt"

// Enums in Go are typically implemented using constants and iota.
// iota is a special identifier that simplifies constant definitions that use incrementing numbers.
// They provide a way to define a set of related named values.
// The main reason to use enums is to improve code readability and maintainability by giving meaningful names to sets of related values.
// set of related values means that these values belong to the same category or type, such as days of the week, order statuses, or user roles.

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Shipped
	Delivered
	Cancelled
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status", status)
}

func main() {
	changeOrderStatus(Received)
	changeOrderStatus(Shipped)
	changeOrderStatus(Delivered)
	changeOrderStatus(Cancelled)
}
