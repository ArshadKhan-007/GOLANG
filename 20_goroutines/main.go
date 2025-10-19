package main

import (
	"fmt"
	"time"
)

// Goroutine is a lightweight thread managed by the Go runtime. used to run functions concurrently.

func task() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second * 2) // Here, after every 2 seconds a number will be printed
	}
}

func main() {
	go task()
	fmt.Println("Main Function Continues")
	time.Sleep(time.Second * 10) // Here, main function will wait for 10 seconds before exiting
	fmt.Println("Main Function Ends")
}

//CONCLUSION:
// In this example, the task function runs as a goroutine, allowing the main function to continue executing without waiting for task to complete.
// The main function sleeps for 10 seconds to ensure that the goroutine has enough time to finish its execution before the program exits.
// SO, Goroutines help in acheiving concurrency which means multiple tasks can run simultaneously without blocking each other.

// REAL LIFE USAGE:
// Goroutines are widely used in web servers to handle multiple client requests concurrently, in data processing pipelines to process data in parallel, and in real-time applications like chat servers or gaming servers where responsiveness is crucial.
