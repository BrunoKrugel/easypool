package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Number of goroutines to create
	numGoroutines := 5

	// Launch multiple goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1) // Increment the WaitGroup counter

		go func(id int) {
			defer wg.Done() // Decrement the WaitGroup counter when done

			// Simulate some work
			fmt.Printf("Goroutine %d started\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Goroutine %d completed\n", id)
		}(i) // Pass 'i' as an argument to the goroutine function
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// All goroutines have completed
	fmt.Println("All goroutines have finished.")
}
