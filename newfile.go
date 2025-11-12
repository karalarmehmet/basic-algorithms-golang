package main

import "fmt"

func main() {

	a := "Şaralar"
	fmt.Println(a[0])

	b := "Şaralarr"
	runes := []rune(b)
	fmt.Println(runes[0])

	// var a string
	// fmt.Println("Enter a string to be reversed")
	// fmt.Scanln(&a)
	// reverse(a)
}

// func reverse(s string) string {
// 	runes := []rune(s)
// 	len := len(runes)
// 	for i := 0; i < len/2; i++ {
// 		runes[i], runes[len-1-i] = runes[len-1-i], runes[i]
// 	}
// 	return string(runes)
// }
