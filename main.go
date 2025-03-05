package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/schollz/progressbar/v3"
)

// Default values used for benchmark
const (
	defaultNumTasks   = 200
	defaultNumWorkers = 30
	defaultFibMin     = 40
	defaultFibMax     = 45
)

// Global progress bar used by several functions
var bar *progressbar.ProgressBar

func main() {
	// Create main argument parser
	parser := argparse.NewParser("Fib-dot-go", "CPU benchmark that calculates Fibonacci numbers")

	// Create command parsers
	runParser := parser.NewCommand("run", "Run benchmark with settable options")
	aboutParser := parser.NewCommand("about", "Display about/help information")
	cpuParser := parser.NewCommand("cpu", "Display CPU information and utilisation")

	// Create run arguments
	runNumTasks := runParser.Int("t", "tasks", &argparse.Options{
		Default: defaultNumTasks,
		Help:    "Total number of Fibonacci numbers to calculate",
	})
	runNumWorkers := runParser.Int("w", "workers", &argparse.Options{
		Default: defaultNumWorkers,
		Help:    "Number of worker threads to spawn",
	})
	runFibMin := runParser.Int("n", "min", &argparse.Options{
		Default: defaultFibMin,
		Help:    "Minimum Fibonacci number to calculate (inclusive)",
	})
	runFibMax := runParser.Int("x", "max", &argparse.Options{
		Default: defaultFibMax,
		Help:    "Maximum Fibonacci number to calculate (exclusive)",
	})
	runMinimal := runParser.Flag("m", "minimal", &argparse.Options{
		Help: "Only display completion time",
	})

	// Create CPU argument
	cpuMinimal := cpuParser.Flag("m", "minimal", &argparse.Options{
		Help: "Only display CPU name and core-count",
	})

	// Parse cli input
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	// Select by command
	if runParser.Happened() {
		// Check validity of cli values
		if *runNumTasks <= 0 {
			fmt.Print(parser.Usage(fmt.Errorf("[-t|--tasks] value must be greater than 0 [%d]", *runNumTasks)))
			return
		}
		if *runNumWorkers <= 0 {
			fmt.Print(parser.Usage(fmt.Errorf("[-w|--workers] value must be greater than 0 [%d]", *runNumWorkers)))
			return
		}
		if *runFibMin >= *runFibMax {
			fmt.Print(parser.Usage(fmt.Errorf("[-n|--min] [-x|--max] min must be less than max [%d] [%d]", *runFibMin, *runFibMax)))
			return
		}

		if !*runMinimal {
			// Print minimal cpu information
			showCPU(true)

			// Create progress bar if not minimal
			bar = progressbar.NewOptions(*runNumTasks,
				progressbar.OptionShowCount(),
				progressbar.OptionSetWidth(60),
				progressbar.OptionSetRenderBlankState(true),
				progressbar.OptionSetTheme(progressbar.Theme{
					Saucer:        "=",
					SaucerHead:    ">",
					SaucerPadding: " ",
					BarStart:      "[",
					BarEnd:        "]",
				}),
			)
		}

		// Run benchmark
		run(*runNumTasks, *runNumWorkers, *runFibMin, *runFibMax, *runMinimal)
	} else if aboutParser.Happened() {
		showAbout()
	} else if cpuParser.Happened() {
		showCPU(*cpuMinimal)
	}
}
