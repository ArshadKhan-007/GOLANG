package main

import "fmt"

func main() {
	//Slices is dynamic array
	var nums []int           //Uninitialized slice is nil
	fmt.Println(nums == nil) //true
	fmt.Println(len(nums))

	var vals = make([]int, 5) //make function is used to create slice with length 5 and capacity 5 becuase capacity is optional it will take length as capacity if you don't pass the capacity
	fmt.Println(cap(vals))    //cap function returns capacity of slice
	fmt.Println(vals)
	fmt.Println(vals == nil) //false

	var elements = make([]int, 5, 10) //make function with length and capacity
	fmt.Println(cap(elements))        // Here capacity is 10

	elements = append(elements, 1) //Appending element to slice
	elements = append(elements, 2)
	elements = append(elements, 3)
	elements = append(elements, 4)
	elements = append(elements, 5)
	elements = append(elements, 6)
	elements = append(elements, 7)
	elements = append(elements, 8)
	elements = append(elements, 9)
	elements = append(elements, 10)
	elements = append(elements, 11) //When we append 11th element capacity is increased to double of current capacity
	elements[0] = 100               //Modifying element at index 0
	fmt.Println(elements)
	fmt.Println(cap(elements)) //Capacity is increased to 10*2=20

	//Copy Function
	var values = make([]int, 0, 10)
	values = append(values, 1) //Appending elements to values slice
	var source = make([]int, len(values))

	copy(source, values) //Copying elements from elements slice to source slice
	fmt.Println(values, source)

	//Slice Operations
	var data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(data[0:2]) //Slicing from index 0 to 2 (2 is exclusive)

	var Data = []int{1, 2, 3, 4, 5}
	fmt.Println(Data[:1]) //Slicing from start to index 1 (1 is exclusive)
	fmt.Println(Data[1:]) //Slicing from index 1 to end
	fmt.Println(Data[:])  //Slicing the whole slice

}
