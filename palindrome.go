package main

import (
	"fmt"
	"strings"
)

func isPalindrome(n string) bool {
	n = strings.ToLower(n)

	reversed := ""

	for i := len(n) - 1; i >= 0; i-- {
		reversed += string(n[i])
	}
	return reversed == n

}

func main() {
	var word string
	fmt.Println("Please enter the string that you want to check wheter it is palindrome or not")
	fmt.Scan(&word)

	if isPalindrome(word) {
		fmt.Printf("%s is a palindrome ", word)
	} else {
		fmt.Printf("%s is not a palindrome", word)
	}

}
