package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Print("please enter a number ")
	fmt.Scan(&num)

	if isPrime(num) {
		fmt.Printf("%d is a prime number", num)
	} else {
		fmt.Printf("%d is not a prime number", num)
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	for i := 5; i <= int(math.Sqrt(float64(n))); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
