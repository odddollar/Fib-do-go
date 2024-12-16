package main

import (
	"fmt"
	"sync"
	"time"
)

// Default values used for benchmark
const (
	numWorkers = 30
	numTasks   = 200
	fibMin     = 40
	fibMax     = 45
)

func main() {
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

		// Give worker id, task channel, and wait group
		go worker(i, tasks, &wg)
	}

	// Close channel when no more tasks
	close(tasks)

	// Wait for all workers to finish
	wg.Wait()

	// Print elapsed time
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)

	// Wait for keypress to exit
	fmt.Print("Press ENTER to exit")
	fmt.Scanln()
}
