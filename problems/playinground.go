package main

import "fmt"

func main() {
	// Our list of colors
	colors := []string{"red", "blue", "red", "green", "blue"}

	// Make a map to store counts
	colorCount := make(map[string]int)

	// Loop through every color in the list
	for _, color := range colors {
		colorCount[color]++ // add 1 each time we see this color
	}

	// Now print the results
	for key, value := range colorCount {
		fmt.Printf("%s: %d\n", key, value)
	}
}
