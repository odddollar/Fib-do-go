package main

import (
	"fmt"
	"sync"
	"time"
)

// Run benchmark with given parameters
func run(numTasks, numWorkers, fibMin, fibMax int, minimal bool) {
	// Save start time
	start := time.Now()

	// Create buffered channel to hold tasks. Acts like a queue
	tasks := make(chan int, numTasks)

	// Create wait group to keep running until all tasks complete
	var wg sync.WaitGroup

	// Push tasks onto queue
	for i := 0; i < numTasks; i++ {
		tasks <- fibMin + (i % (fibMax - fibMin))
	}

	// Spawn workers
	for i := 1; i <= numWorkers; i++ {
		// Increment wait group
		wg.Add(1)

		// Give worker task channel, wait group, and progress bar
		go worker(tasks, &wg)
	}

	// Close channel when no more tasks
	close(tasks)

	// Wait for all workers to finish
	wg.Wait()

	// Print elapsed time
	elapsed := time.Since(start)
	if minimal {
		fmt.Printf("%.3f", elapsed.Seconds())
	} else {
		fmt.Printf("\nTime taken: %.3fs\n\n", elapsed.Seconds())
	}
}
