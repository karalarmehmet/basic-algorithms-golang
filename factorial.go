package main

import (
	"fmt"
)

func factorialRecursion(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * factorialRecursion(n-1)
	}
}

func main() {
	var input int
	fmt.Print("Please enter a number: ")
	for {
		// Read an integer input directly
		_, err := fmt.Scan(&input)
		if err != nil {
			// If input is not a valid integer
			fmt.Println("Please enter a valid number")

			var dummy string
			fmt.Scanln(&dummy)
			continue

		}

		// Check for negative numbers
		if input < 0 {
			fmt.Println("Factorial cannot be calculated for negative numbers")
			return
		}

		// Calculate factorial
		result := factorialRecursion(input)
		fmt.Printf("The result of %d! is %d\n", input, result)
		return

	}
}
