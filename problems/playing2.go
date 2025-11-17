package main

import "fmt"

func main() {
	var r rune = 'A' // value is 65, type int32
	var i int = 10   // Type int (maybe int64)

	// ERROR: invalid operation: r + i (mismatched types rune and int)
	// toplam1 := r + i
	// fmt.Println(toplam1)
	sum := int(r) + i
	fmt.Println(sum)

	res := rune(sum)
	fmt.Println(res)

	fmt.Println("%d\n", res)
}
