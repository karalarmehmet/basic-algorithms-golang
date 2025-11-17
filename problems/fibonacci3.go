package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter fibonacci number that you want to calculate")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)

	if number < 0 || err != nil {
		fmt.Println("Enter a valid number")
	}

	fmt.Println(fibonacci(number))
}

func fibonacci(a int) int {
	if a <= 1 {
		return a
	}
	prev := 0
	curr := 1

	for i := 2; i <= a; i++ {
		prev, curr = curr, curr+prev

	}
	return curr
}

// When you call reader.ReadString('\n'), Go reads everything up until the newline and stores it in the input variable.
// The internal buffer that bufio.NewReader uses is automatically managed. After ReadString reads the input, it clears the buffer for the next read operation. So, the buffer is cleared after each call to ReadString, whether the input is valid or not.
// If you enter -4, the input "-4\n" is read into the buffer and stored in input. The buffer is then cleared, and the next time the program asks for input, it is ready to take new data without any leftover characters.
