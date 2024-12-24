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
		Help:    "Minimum fibonacci number to calculate (inclusive)",
	})
	fibMax := runParser.Int("x", "max", &argparse.Options{
		Default: defaultFibMax,
		Help:    "Maximum fibonacci number to calculate (exclusive)",
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
		if *numTasks <= 0 {
			fmt.Print(parser.Usage(fmt.Errorf("[-t|--tasks] value must be greater than 0 [%d]", *numTasks)))
			return
		}
		if *numWorkers <= 0 {
			fmt.Print(parser.Usage(fmt.Errorf("[-w|--workers] value must be greater than 0 [%d]", *numWorkers)))
			return
		}
		if *fibMin >= *fibMax {
			fmt.Print(parser.Usage(fmt.Errorf("[-n|--min] [-x|--max] min must be less than max [%d] [%d]", *fibMin, *fibMax)))
			return
		}

		// Create progress bar
		bar := progressbar.NewOptions(*numTasks,
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
			progressbar.OptionOnCompletion(func() { fmt.Println() }),
		)

		// Run benchmark
		run(*numTasks, *numWorkers, *fibMin, *fibMax, bar)
	} else if aboutParser.Happened() {
		showAbout()
	}
}
