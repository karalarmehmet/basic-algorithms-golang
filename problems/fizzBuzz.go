package main

import "fmt"

func main() {
	var a int
	fmt.Println("Enter a number to test fizz buzz")
	fmt.Scan(&a)
	fizzbuzz(a)
}
func fizzbuzz(a int) {
	for i := 1; i <= a; i++ {
		switch {
		case i%5 == 0 && i%3 == 0:
			fmt.Print("fizzbuzz")
		case i%3 == 0:
			fmt.Print("fizz ")
		case i%5 == 0:
			fmt.Print("buzz ")
		default:
			fmt.Print(i)
		}
		if i != a {
			fmt.Print(",")
		}
	}
}
