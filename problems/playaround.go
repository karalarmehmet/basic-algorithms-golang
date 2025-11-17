package main

import "fmt"

func main() {
	var myByte byte = 65
	var myRune rune = 65
	var myInt int = 65
	var myString string = "mehmet"

	fmt.Printf("this is the number %d and the type for myByte:%T\n", myByte, myByte)
	fmt.Printf("this is the number %d and the type for myRune:%T\n", myRune, myRune)
	fmt.Printf("this is the number %d and the type for myInt:%T\n", myInt, myInt)
	fmt.Printf("this is the string %s and the type for myString:%T\n", myString, myString)
}
