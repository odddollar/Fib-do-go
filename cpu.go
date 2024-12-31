package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
)

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
	}
	fmt.Println("Cores:", physical)

	// Get CPU threads
	logical, err := cpu.Counts(true)
	if err != nil {
		fmt.Println("Error retrieving CPU thread-count:", err)
	}
	fmt.Println("Threads:", logical)
}
