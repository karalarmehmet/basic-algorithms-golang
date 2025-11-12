package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// a   b   a    b   b    b
// 0   1   1    2   3    5   8   13   21   34
//	   a   b    a

func fibonacciIterative(n int) {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b

	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the number of fibonacci terms ")

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)

		if err != nil || num < 0 {
			fmt.Println("Please enter a valid positvie number")

			continue

		}
		fmt.Printf("First %d fibonacci numbers ", num)
		fibonacciIterative(num)
		break
	}

}

// When you call reader.ReadString('\n'), Go reads everything up until the newline and stores it in the input variable.
// The internal buffer that bufio.NewReader uses is automatically managed. After ReadString reads the input, it clears the buffer for the next read operation. So, the buffer is cleared after each call to ReadString, whether the input is valid or not.
// If you enter -4, the input "-4\n" is read into the buffer and stored in input. The buffer is then cleared, and the next time the program asks for input, it is ready to take new data without any leftover characters.
