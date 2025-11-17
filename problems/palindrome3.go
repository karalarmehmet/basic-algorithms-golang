package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter a string that you want to chaeck whether it's a palindrome or nor ")

	for {
		str, err := reader.ReadString('\n')
		str = strings.TrimSpace(str)

		if err != nil {
			fmt.Println("please enter something valid!")

		}
		if palindrome(str) {
			fmt.Printf("%s is a palindrome ", str)
			break

		} else {
			fmt.Printf("%s is not a palindrome ", str)
			break
		}

	}

}

func palindrome(n string) bool {

	length := len(n)

	for i := 0; i < length/2; i++ {
		if n[i] != n[length-1-i] {
			return false
		}

	}
	return true
}
