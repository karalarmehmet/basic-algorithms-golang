package main

import "fmt"

func main() {
	fmt.Println("input 2 numbers for GCD")
	var a int
	var b int
	fmt.Scan(&a, &b)
	fmt.Println(gcd(a, b))
}

// func gcd(a, b int) int {
// 	for b != 0 {
// 		a, b = b, a%b
// 	}
// 	return a
// }

func gcd(a, b int) int {
	var limit int
	var gcd int
	if a > b {
		limit = b
	} else {
		limit = a
	}
	for i := limit; i > 0; i-- {
		if a%i == 0 && b%i == 0 {
			gcd = i //OR return i
			break
		}

	}
	return gcd // return 1 (if we say return i in the for loop then we return 1) normally the function never hits here just because we have return in the function definition we need to return something
}
