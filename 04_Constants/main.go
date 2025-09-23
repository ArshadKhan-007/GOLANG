package main

import "fmt"

func main() {
	//Constant
	const name = "golang"
	// name = "python" //Cannot change the value of a constant
	fmt.Println(name)

	const (
		port = 8080
		host = "localhost"
	)
	fmt.Println(port, host)
}
