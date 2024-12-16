package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

// Default values used for benchmark
const (
	defaultNumTasks   = 200
	defaultNumWorkers = 30
	defaultFibMin     = 40
	defaultFibMax     = 45
)

func main() {
	// Create argument parser
	parser := argparse.NewParser("fib-dot-go", "CPU benchmark that calculates fibonacci numbers")

	// Create arguments
	numTasks := parser.Int("t", "tasks", &argparse.Options{Default: defaultNumTasks})
	numWorkers := parser.Int("w", "workers", &argparse.Options{Default: defaultNumWorkers})
	fibMin := parser.Int("m", "min", &argparse.Options{Default: defaultFibMin})
	fibMax := parser.Int("x", "max", &argparse.Options{Default: defaultFibMax})

	// Parse cli input
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	// Run benchmark
	run(*numTasks, *numWorkers, *fibMin, *fibMax)
}
