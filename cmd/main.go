package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/BrunoKrugel/easypool"
)

func main() {
	worker := easypool.NewPool(5)

	var wg sync.WaitGroup

	//Number of routines to be executed
	for i := 0; i < 10; i++ {
		taskNumber := i
		wg.Add(1)
		worker.AddToPool(func() {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Printf("Task %d executed\n", taskNumber)
		})
	}

	worker.Execute()
	wg.Wait()

	worker.Close()

	fmt.Println("All goroutines have finished.")
}
