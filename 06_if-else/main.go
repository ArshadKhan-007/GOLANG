package main

import "fmt"

func main() {
	age := 18
	if age >= 18 {
		fmt.Println("Eligible")
	} else {
		fmt.Println("Not Eligible")
	}

	role := "admin"
	hasPermission := false
	if role == "admin" && hasPermission {
		fmt.Println("Access Granted")
	} else {
		fmt.Println("Access Denied")
	}

	Age := 16
	if Age >= 18 {
		fmt.Println("An Adult")
	} else if Age >= 16 {
		fmt.Println("A Teenager")
	} else {
		fmt.Println("A Child")
	}

	//Declaring a variable inside if statement
	if Age := 16; Age >= 18 {
		fmt.Println("An Adult")
	} else if Age >= 16 {
		fmt.Println("A Teenager")
	} else {
		fmt.Println("A Child")
	}

}
