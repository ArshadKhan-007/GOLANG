package main

import (
	"fmt"
	"time"
)

func main() {
	Age := 16
	switch {
	case Age >= 18:
		fmt.Println("An Adult")
	case Age >= 16:
		fmt.Println("A Teenager")
	default:
		fmt.Println("A Child")
	}

	i := 2
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Other")
	}

	//Multiple condition in switch
	switch i {
	case 1, 2, 3:
		fmt.Println("One, Two or Three")
	default:
		fmt.Println("Other")
	}

	switch time.Now().Weekday() {
	case time.Wednesday, time.Thursday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}
}
