package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a base and string that needs to be converted to decimal")
	var a int
	var b string
	fmt.Scanf("%d %s", &a, &b)
	res := baseToDec(a, b)
	fmt.Println(res)
}

func baseToDec(a int, b string) int {
	const digits = "0123456789ABCDEF"
	runes := []rune(b)
	var res int
	power := 1

	for i := len(runes) - 1; i >= 0; i-- {
		index := -1
		for j, char := range digits {
			if char == runes[i] {
				index = j
			}
		}
		if index < 0 {
			panic("something went wrong")
		}

		res = res + index*power
		power = power * a
	}
	// for i := len(runes) - 1; i >= 0; i-- {
	// 	res = res + (int(runes[i])-int('0'))*power
	// 	power = power * a
	// }
	return res
}
