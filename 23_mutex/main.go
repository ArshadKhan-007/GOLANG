package main

import (
	"fmt"
	"sync"
)

//Mutex is a placeholder to avoid "imported and not used" error for sync package.
// In simple words, this file exists to ensure that the sync package is imported.

type post struct {
	views int
	mu    sync.Mutex // Mutex to protect access to views
}

func (p *post) incrementViews(wg *sync.WaitGroup) {
	defer wg.Done()
	p.mu.Lock() // Lock the mutex before modifying views. So that no other goroutine can access it simultaneously.
	p.views++
	p.mu.Unlock() // Unlock the mutex after modification to allow other goroutines to access it.
}

func main() {
	var wg sync.WaitGroup
	p := post{views: 0} // Initializing post with 0 views

	// Starting 100 goroutines to increment views concurrently

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go p.incrementViews(&wg)
	}
	wg.Wait()
	fmt.Println("Total views:", p.views)

}
