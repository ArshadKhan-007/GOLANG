package main

import (
	"fmt"
	"sync"
)

// WaitGroups are used to wait for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of goroutines to wait for.
// Then each of the goroutines runs and calls Done when finished.
// At the same time, Wait can be used to block until all goroutines have finished.

// WaitGroups is introduced bcz of some cons of using sleep method to wait for goroutines to finish their execution.
// Sleep method is not reliable as we can't be sure how much time a goroutine will take to finish its execution.

func task(id int, wg *sync.WaitGroup) { // Here, we have to pass the pointer bcz we want to modify the original WaitGroup
	defer wg.Done() // Decrement the counter when the goroutine completes bcz we called Add before starting the goroutine so it returns to original state after the goroutine is done.
	// Simulating some work with Println
	fmt.Println("Doing task", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1) // Increment the counter for each goroutine we are going to start. It will add according loop we have here.
		// Start a goroutine
		go task(i, &wg) // Passing the address of wg to the goroutine so that it can call Done on the same WaitGroup instance.
	}

	wg.Wait() // Block until the counter is zero, i.e., all goroutines have called Done.
	fmt.Println("All tasks completed.")
}
