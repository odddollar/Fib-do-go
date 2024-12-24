package main

import (
	"sync"
)

// Worker pulls next task off tasks queue and executes. Finishes when
// no more tasks
func worker(tasks <-chan int, wg *sync.WaitGroup) {
	// Complete when no more tasks
	defer wg.Done()

	// Pull next task off queue
	for task := range tasks {
		_ = fib(task)

		// Update progress bar
		if bar != nil {
			bar.Add(1)
		}
	}
}
