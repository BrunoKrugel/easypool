package easypool

import (
	"sync"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	t.Run("Test WorkerPool", func(t *testing.T) {
		poolSize := 3
		wp := NewPool(poolSize)
		var wg sync.WaitGroup

		for i := 0; i < poolSize*2; i++ {
			taskNumber := i
			wg.Add(1)
			wp.AddToPool(func() {
				defer wg.Done()
				time.Sleep(time.Millisecond)
				t.Logf("Task %d executed", taskNumber)
			})
		}
		wp.Execute()
		wg.Wait()
		wp.Close()
	})
}
