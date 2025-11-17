package main

import "fmt"

func main() {
	var a int
	var b string
	var t int
	charMap := map[rune]int{
		'A': 10,
		'B': 11,
		'C': 12,
		'D': 13,
		'E': 14,
		'F': 15,
	}
	fmt.Println("please enter a base it's value and a target base")
	fmt.Scanf("%d %s %d ", &a, &b, &t)
	dec := baseToDec(a, b)
	base := DecToBase(dec, t, charMap)
	fmt.Println(base)
}
