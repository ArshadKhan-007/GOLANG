package main

import (
	"fmt"
	"time"
)

// Structs in Go are used to group related data together to form a single entity. They are similar to classes in other programming languages but do not have methods associated with them.

// A struct is defined using the `type` keyword followed by the name of the struct and the `struct` keyword. The fields of the struct are defined within curly braces.

type order struct {
	orderId   int
	amount    float64
	status    string
	createdAt time.Time
}

// Method to change the amount of the order
// In Go, methods can be defined on struct types. Here, we define a method `changeAmount` that takes a float64 parameter and updates the `amount` field of the `order` struct.
// The receiver `o *order` indicates that this method is associated with the `order` struct and can modify its fields.
// The receiver is a pointer to the struct, allowing the method to modify the original struct instance.
func (o *order) changeAmount(amount float64) {
	o.amount = amount // Here, we havn't used *o to access the amount field because Go automatically dereferences the pointer when accessing struct fields.
}

func main() {
	// Creating an instance of the struct
	// instance is similar to object in other languages, where we have to define object for a class to use it.
	// Here, we are creating an instance of the `order` struct and initializing its fields.
	o := order{
		orderId:   1,
		amount:    99.99,
		status:    "Pending",
		createdAt: time.Now(),
	} // we can have multiple instances of the same struct with different values.
	// Make sure that we must provide different names for different structs instances.

	o.changeAmount(999.99) // Here, we are calling the `changeAmount` method on the `order` struct instance `o` to update its `amount` field.

	// Changing struct fields
	o.status = "Completed" // Here, we are updating the `status` field of the `order` struct instance `o`.

	// Accessing struct fields
	fmt.Printf("Order ID: %d\n", o.orderId)                          // Here, %d is used to format integer values.
	fmt.Printf("Amount: %.2f\n", o.amount)                           // Here, %.2f is used to format float values to 2 decimal places.
	fmt.Printf("Status: %s\n", o.status)                             // Here, %s is used to format string values.
	fmt.Printf("Created At: %s\n", o.createdAt.Format(time.RFC1123)) // Here, we format the time using RFC1123 format.
}
