package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Println("plaese enter a string that you want to reverse")
	n, _ := fmt.Scan(&s)

	reverseString(s)
	fmt.Println(n)
}

func reverseString(n string) string {

	reversed := ""

	for i := len(n) - 1; i >= 0; i-- {
		reversed += string(n[i])
	}
	return reversed

}
