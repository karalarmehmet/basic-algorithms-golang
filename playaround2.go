package main

import "fmt"

func main() {
	// var a rune = 'A'
	// fmt.Printf("%c %d\n", a, a)
	// var char1 byte = 'A'
	// fmt.Printf("this is the literal %d \n ", char1)
	// var b byte = 'ü'
	// fmt.Printf("the number %d", b)

	// var r rune = 'A' // value 65, type int32
	// // var i int = 10   // value 10 ,type is maybe int64 or int32 depends on the architecture
	// fmt.Printf("%U %d %c\n", r, r, r)

	// var a rune = 'ü'
	r := 'ü'       // U+00FC
	s := string(r) // UTF-8: 0xC3 0xBC
	fmt.Printf("%U", r)
	fmt.Println(s)

}
