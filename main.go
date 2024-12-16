package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Save start time
	start := time.Now()

	// Constants to adjust benchmark
	const numWorkers = 30
	const numTasks = 200
	const fibMin = 40
	const fibMax = 45
	const fibDiff = fibMax - fibMin

	// Create buffered channel to hold tasks. Acts like a queue
	tasks := make(chan int, numTasks)

	// Create wait group to keep running until all tasks complete
	var wg sync.WaitGroup

	// Push tasks onto queue
	for i := 0; i < numTasks; i++ {
		tasks <- fibMin + (i % fibDiff)
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
