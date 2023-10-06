package easypool

import (
	"sync"
)

type WorkerPool struct {
	poolSize int
	queue    chan func()
	wg       sync.WaitGroup
	done     chan struct{}
}

// NewPool creates a new WorkerPool with the specified size.
func NewPool(poolSize int) *WorkerPool {
	return &WorkerPool{
		poolSize: poolSize,
		queue:    make(chan func(), poolSize),
		done:     make(chan struct{}),
	}
}

// AddToPool adds a function to be executed by the worker pool.
func (wp *WorkerPool) AddToPool(task func()) {
	wp.wg.Add(1)
	go func() {
		defer wp.wg.Done()
		wp.queue <- task
	}()
}

// Launch the worker pool and executes the functions in the pool.
func (wp *WorkerPool) Execute() {
	for i := 0; i < wp.poolSize; i++ {
		wp.wg.Add(1)
		go func() {
			defer wp.wg.Done()
			for {
				select {
				case task, ok := <-wp.queue:
					if !ok {
						return
					}
					task()
				case <-wp.done:
					return
				}
			}
		}()
	}
}

// Waits for all tasks in the worker pool to complete.
func (wp *WorkerPool) Wait() {
	close(wp.queue)
	wp.wg.Wait()
}

// Closes the worker pool gracefully.
func (wp *WorkerPool) Close() {
	close(wp.done)
}
