package main

import "fmt"

//This changeNum function is here to state without pointer the value of num in main will not change
// It takes an integer as an argument and attempts to change its value to 5
// However, this change does not affect the original variable in main
// because integers are passed by value in Go
func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum", num)
}

// Pass by reference using pointers
func changeNumByRef(val *int) {
	*val = 5 // Dereference the pointer and change the value at that address to 5
	fmt.Println("In changeNumByRef", *val)
}

func main() {
	num := 1       // initial value of num is 1
	changeNum(num) // call changeNum with num
	// After calling changeNum, the value of num in main remains unchanged
	fmt.Println("After change in main", num)

	val := 1             // initial value of val is 1
	changeNumByRef(&val) // call changeNumByRef with the address of val
	// After calling changeNumByRef, the value of val in main is changed to 5
	fmt.Println("After change in main", val)

	// As you can see, here we haven't used * symbol
	// Why we haven't use * symbol because we are not dereferencing the pointer
	// We are just assigning the address of value to ptr
	//
	value := 10
	ptr := &value
	fmt.Println("Value:", value) // prints 10
	fmt.Println("Pointer:", ptr) // prints the memory address of value

	// Here, we have used * symbol to dereference the pointer and get the value at that address
	//
	var valStr string
	valStr = "Hello"
	var ptrstr *string              // Here * indicates that ptrstr is a pointer to a string type
	ptrstr = &valStr                // Assign the address of valStr to ptrstr
	fmt.Println("Value:", valStr)   // prints "Hello"
	fmt.Println("Pointer:", ptrstr) // prints the memory address of valStr

}
