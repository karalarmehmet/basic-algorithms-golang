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

func fibonacciIterative(n int) []int {
	res := make([]int, n+1) //	res := make([]int, 0, n+1)
	a, b := 0, 1
	for i := 0; i < n; i++ {
		res[i] = a // if we use "res := make([]int, 0, n+1)" then in this line we use res=append(res,a) but now we assigne indexes
		a, b = b, a+b
	}
	return res
}

func main() {
	fmt.Println("Enter the number of fibonacci terms")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)

	if number < 0 || err != nil {
		fmt.Println("please enter a valid number of terms")
	}
	res := fibonacciIterative(number)
	fmt.Println(res)
}

// When you call reader.ReadString('\n'), Go reads everything up until the newline and stores it in the input variable.
// The internal buffer that bufio.NewReader uses is automatically managed. After ReadString reads the input, it clears the buffer for the next read operation. So, the buffer is cleared after each call to ReadString, whether the input is valid or not.
// If you enter -4, the input "-4\n" is read into the buffer and stored in input. The buffer is then cleared, and the next time the program asks for input, it is ready to take new data without any leftover characters.
