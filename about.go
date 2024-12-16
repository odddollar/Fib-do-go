package main

import (
	"fmt"

	"github.com/mitchellh/go-wordwrap"
)

// Print about information
func showAbout() {
	aboutString := "Fib-dot-go is a CPU benchmark that aims to assess performance in " +
		"calculating fibonacci numbers. There is a set queue of tasks that " +
		"worker threads pull from, where each task consists of recursively " +
		"calculating a large fibonacci number.\n\nFor example, with a minimum " +
		"fibonacci number of 40, a maximum of 45, and 200 tasks, calculating " +
		"the 40th fibonacci number is pushed onto a queue. Calculating the " +
		"41st number is pushed next, then the 42nd, etc., until the 44th " +
		"number is pushed. The 40th number is then pushed again and the loop " +
		"repeats until a total of 200 tasks exist in the queue. The queue " +
		"would appear as [40 41 42 43 44 40 41 42 43 ...].\n"

	// Wrap and print text
	fmt.Println(wordwrap.WrapString(aboutString, 70))
}
