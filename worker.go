package main

import (
	"sync"

	"github.com/schollz/progressbar/v3"
)

// Worker pulls next task off tasks queue and executes. Finishes when
// no more tasks
func worker(tasks <-chan int, wg *sync.WaitGroup, bar *progressbar.ProgressBar) {
	// Complete when no more tasks
	defer wg.Done()

	// Pull next task off queue
	for task := range tasks {
		_ = fib(task)

		if bar != nil {
			bar.Add(1)
		}
	}
}
