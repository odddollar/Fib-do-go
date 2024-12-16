package main

import (
	"fmt"
	"sync"
)

// Worker pulls next task off tasks queue and executes. Finishes when
// no more tasks
func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	// Complete when no more tasks
	defer wg.Done()

	// Pull next task off queue
	for task := range tasks {
		fmt.Printf("Worker %d performing task %d\n", id, task)
		_ = fib(task)
		fmt.Printf("Worker %d completed task %d\n", id, task)
	}
}
