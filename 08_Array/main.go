package main

import "fmt"

func main() {
	//Array
	var nums [5]int

	//Assigning values to array elements
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	nums[3] = 40
	nums[4] = 50

	fmt.Println(nums) //Printing the entire array

	fmt.Println(nums[0]) //Accessing array elements using index

	fmt.Println(len(nums)) //len function returns the length of the array

	num := [4]int{1, 2, 3, 4} //Array initialization with values
	fmt.Println(num)

	// 2-D Array
	matrix := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(matrix)

}
