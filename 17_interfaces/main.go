package main

import "fmt"

// Interface are a way to define behavior in Go. They specify a set of method signatures that a type must implement to satisfy the interface.
// Use interface to achieve polymorphism, allowing different types to be treated uniformly based on shared behavior.

// Interface 1
type Reader interface {
	Read() string
}

// Interface 2
type writer interface {
	Write(data string)
}

// Interface 3: ReadWriter combines Reader and writer interfaces
type ReadWriter interface {
	Reader
	writer
}

// File struct implements the ReadWriter interface
type File struct {
	content string
}

// Implementing Read method for File
func (f *File) Read() string {
	return f.content
}

// Implementing Write method for File
func (f *File) Write(data string) {
	f.content = data
	fmt.Println("Data written to file.")
}

// Function that takes a ReadWriter interface as a parameter
func process(rw ReadWriter) {
	fmt.Println("Processing file....")
	rw.Write("Go interfaces are lit")
	fmt.Println("Reading back:", rw.Read())
}

func main() {
	// Creating an instance of File and passing it to process function
	file := &File{}
	process(file)
}
