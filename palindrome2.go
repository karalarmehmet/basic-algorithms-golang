package main

//palindrome only for numbers
import "fmt"

func isPalindrome2(p int) bool {

	original := p
	reversed := 0

	for p > 0 {
		digit := p % 10
		reversed = reversed*10 + digit
		p /= 10
	}
	return original == reversed
}

func main() {
	var a int
	fmt.Println("please enter to check palindrome number")

	fmt.Scan(&a)

	if isPalindrome2(a) {
		fmt.Printf("%d is a palindrome number", a)
	} else {
		fmt.Printf("%d is not a palindrome number", a)
	}

}
