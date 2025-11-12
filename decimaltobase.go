package main

import (
	"fmt"
	"strconv"
)

// decimal is base 10
func main() {
	var a int
	var b int
	fmt.Println("Enter a Dec and a Base for converter")
	fmt.Scanf("%d %d", &a, &b)
	charMap := map[rune]int{
		'A': 10,
		'B': 11,
		'C': 12,
		'D': 13,
		'E': 14,
		'F': 15,
	}
	res := DecToBase(a, b, charMap)
	fmt.Println(res)

}

func DecToBase(a, b int, c map[rune]int) string {
	var reminder int
	var res string

	for a > 0 {
		reminder = a % b
		a = a / b
		if reminder < 10 {
			res = strconv.Itoa(reminder) + res
		} else if reminder >= 10 {
			for i, val := range c {
				if reminder == val {
					res = string(i) + res
				}
			}
		}
	}
	return res
}

// func main() {
// 	fmt.Println("enter a dec and a base for converter")
// 	var a, b int
// 	fmt.Scanf("%d %d", &a, &b)
// 	res := DecToBase(a, b)
// 	fmt.Println(res)
// }
// func DecToBase(dec, base int) string {
// 	const digits = "0123456789ABCDEF"
// 	var sb strings.Builder

// 	for dec > 0 {
// 		rem := dec % base
// 		sb.WriteByte(digits[rem])
// 		dec = dec / base

// 	}
// 	result := sb.String()
// 	runes := []rune(result)

// 	for i := 0; i <= len(runes)/2; i++ {				alternatif:	for i,j:=0,len(runes)-1;i<j;i,k=i+1,j-1
// 		for j := len(runes) - 1; j > i; j-- {						 runes[i], runes[j] = runes[j], runes[i]
// 			if j == len(runes)-i-1 {
// 				runes[i], runes[j] = runes[j], runes[i]
// 				break
// 			}

// 		}
// 	}
// 	return string(runes)
// }
