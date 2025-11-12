package main

import "fmt"

func main() {
	res := romanToInt("XLIII")
	fmt.Println(res)
}
func romanToInt(s string) int {
	romMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	prev := 0
	for _, v := range s {
		cur := romMap[v]
		if cur > prev {
			result += cur-
		} else {
			result = result + romMap[string(s[i])]
			i += 1
		}
		cur = prev
	}

	return result
}
