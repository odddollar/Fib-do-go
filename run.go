package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
)

func run(numTasks, numWorkers, fibMin, fibMax int, bar *progressbar.ProgressBar) {
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
		go worker(tasks, &wg, bar)
	}

	// Close channel when no more tasks
	close(tasks)

	// Wait for all workers to finish
	wg.Wait()

	// Print elapsed time
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %.3fs\n\n", elapsed.Seconds())
}
