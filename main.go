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
	// Create main argument parser
	parser := argparse.NewParser("fib-dot-go", "CPU benchmark that calculates fibonacci numbers")

	// Create command parsers
	runParser := parser.NewCommand("run", "Run benchmark with settable options")
	aboutParser := parser.NewCommand("about", "Display about/help information")

	// Create run arguments
	numTasks := runParser.Int("t", "tasks", &argparse.Options{
		Default: defaultNumTasks,
		Help:    "Total number of fibonacci numbers to calculate",
	})
	numWorkers := runParser.Int("w", "workers", &argparse.Options{
		Default: defaultNumWorkers,
		Help:    "Number of worker threads to spawn",
	})
	fibMin := runParser.Int("n", "min", &argparse.Options{
		Default: defaultFibMin,
		Help:    "Minimum fibonacci number to calculate",
	})
	fibMax := runParser.Int("x", "max", &argparse.Options{
		Default: defaultFibMax,
		Help:    "Maximum fibonacci number to calculate",
	})

	// Parse cli input
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	if runParser.Happened() {
		// Run benchmark
		run(*numTasks, *numWorkers, *fibMin, *fibMax)
	} else if aboutParser.Happened() {
		return
	}

	// Wait for keypress to exit
	fmt.Print("Press ENTER to exit")
	fmt.Scanln()
}
