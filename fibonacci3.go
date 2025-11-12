package main

import "fmt"

// Iterative approach using a slice
func fibonacciIterative1(n int) int {
	if n <= 1 {
		return n
	}
	// Create a slice to store Fibonacci numbers
	fib := make([]int, n+1)
	fib[0], fib[1] = 0, 1

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}

func main() {
	var num int
	fmt.Println("Enter the number of Fibonacci terms")

	for {
		a, err := fmt.Scan(&num) // a is the number of items that is scanned succesfully we dont have anything to do with that
		fmt.Println(a)

		if err != nil || num < 0 {
			fmt.Println("please enter a valid input")
			continue
		}
		break
	}

	// Compute the nth Fibonacci number
	result := fibonacciIterative1(num)
	fmt.Printf("The %dth Fibonacci number is %d\n", num, result)
}
