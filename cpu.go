package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

// Print CPU model, core and thread-count, and usage
func showCPU(minimal bool) {
	// Get CPU model
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Println("Error retrieving CPU model:", err)
		return
	}
	fmt.Println("Model:", cpuInfo[0].ModelName)

	// Get CPU cores
	physical, err := cpu.Counts(false)
	if err != nil {
		fmt.Println("Error retrieving CPU core-count:", err)
		return
	}
	fmt.Println("Cores:", physical)

	// Get CPU threads
	logical, err := cpu.Counts(true)
	if err != nil {
		fmt.Println("Error retrieving CPU thread-count:", err)
		return
	}
	fmt.Println("Threads:", logical)
	fmt.Println()

	// Get CPU usage
	if !minimal {
		for {
			percentages, err := cpu.Percent(1*time.Second, false)
			if err != nil {
				fmt.Println("Error retrieving CPU usage:", err)
				break
			}

			if len(percentages) > 0 {
				fmt.Printf("CPU Usage: %.2f%%\n", percentages[0])
			}
		}
	}
}
