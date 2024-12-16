package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

// Default values used for benchmark
const (
	defaultNumWorkers = 30
	defaultNumTasks   = 200
	defaultFibMin     = 40
	defaultFibMax     = 45
)

func main() {
	// Create argument parser
	parser := argparse.NewParser("fib-dot-go", "CPU benchmark that calculates fibonacci numbers")

	// Parse cli input
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	// Run benchmark
	run(defaultNumWorkers, defaultNumTasks, defaultFibMin, defaultFibMax)

}
