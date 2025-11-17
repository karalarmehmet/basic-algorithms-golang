package main

import "fmt"

func main() {
	// str := "😊a"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Print(str[i])
	// 	if i != len(str) {
	// 		fmt.Print(",")
	// 	}
	// }
	// for _, b := range str {
	// 	fmt.Print(b)
	// 	fmt.Print(",")
	// }
	// fmt.Println()
	// fmt.Println(len(str))
	// byte := []byte(str)
	// fmt.Println(byte)

	var r2 string = "😊ü"
	// fmt.Println(r2[:5]) //slicing returns a new string
	for _, r := range r2 {
		fmt.Println(r)
	}
}
