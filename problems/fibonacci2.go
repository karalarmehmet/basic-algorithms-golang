package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the number of fibonacci terms")

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err != nil || num < 0 {
			fmt.Println("please enter a valid input")
			continue
		}

		for i := 0; i < num; i++ {
			fmt.Print(fibonacciRecursive(i), " ")

		}
		break
	}
}
