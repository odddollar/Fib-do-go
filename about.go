package main

import (
	"fmt"
)

// Print about information
func showAbout() {
	aboutString := `Fib-dot-go is a multithreaded CPU benchmark that aims to assess processor 
performance in calculating fibonacci numbers. There is a queue of tasks that 
worker threads pull from, where each task consists of recursively calculating 
a large fibonacci number.

For example, with a minimum fibonacci number of 40, a maximum of 45, and 200 
tasks, calculating the 40th fibonacci number is pushed onto a queue. Calculating 
the 41st number is pushed next, then the 42nd, etc., until the 44th number is 
pushed. The 40th number is then pushed again and the loop repeats until a total 
of 200 tasks exist in the queue. The queue would appear as [40 41 42 43 44 40 
41 42 43 ...]. Worker threads pull the next available task from this queue and 
calculate that fibonacci number. Once this task has been completed, the worker 
pulls the next task until all tasks have been pulled from the queue. A single 
shared queue is used by all workers.

On exhaustion of all tasks, the time taken to complete the benchmark is reported.

Arguments:
  -t  --tasks    Controls the number of tasks to complete, and thus the overall 
                 length of the benchmark.
  -w  --workers  Controls how multithreaded the benchmark is. Increasing this 
                 decreases the time taken, but only up to a certain point. Increase 
                 this value for higher core-count CPUs to maximise utilisation.
  -n  --min      Controls the minimum difficulty of possible tasks. Increasing this 
                 raises the baseline difficulty of the benchmark.
  -x  --max      Controls the maximum difficulty of possible tasks and the overall 
                 range of task difficulty. A greater difficulty range means that some 
                 workers get significantly easier tasks than others, whilst some get 
                 significantly harder tasks.
`

	// Wrap and print text
	fmt.Println(aboutString)
}
