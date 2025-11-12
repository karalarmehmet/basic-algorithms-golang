package main

import "fmt"

func reverseString2(n int) int {

	reverse := 0

	for n > 0 {
		digit := n % 10 // original stringin son basamağı
		reverse = reverse*10 + digit
		n /= 10

	}
	return reverse

}

func main() {
	var s int

	fmt.Println("please enter a number that you want to reverse")
	fmt.Scan(&s)

	result := reverseString2(s)
	fmt.Println(result)
}
