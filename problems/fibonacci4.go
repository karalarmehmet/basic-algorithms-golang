package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the number of the fibonacci that you want to find")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // "1 2 3 4 5\n" OR "5\n"
	input = strings.TrimSpace(input)    //"1 2 3 4 5"   OR "5"
	number, _ := strconv.Atoi(input)
	res := fibonacci_recursive(number)
	fmt.Println(res)

}

func fibonacci_recursive(a int) int {
	if a < 2 && a >= 0 {
		return a
	}
	return fibonacci_recursive(a-2) + fibonacci_recursive(a-1)

}
