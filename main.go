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

	// Create arguments
	numWorkers := parser.Int("w", "workers", &argparse.Options{Default: defaultNumWorkers})
	numTasks := parser.Int("t", "tasks", &argparse.Options{Default: defaultNumTasks})
	fibMin := parser.Int("m", "min", &argparse.Options{Default: defaultFibMin})
	fibMax := parser.Int("x", "max", &argparse.Options{Default: defaultFibMax})

	// Parse cli input
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	// Run benchmark
	run(*numWorkers, *numTasks, *fibMin, *fibMax)
}
