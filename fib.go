package main

// Calculate nth fibonacci number
func fib(n int) int {
	// Base case
	if n <= 1 {
		return n
	}

	// Recursively calculate number
	return fib(n-1) + fib(n-2)
}
