package main

import "fmt"

func main() {
	s := "Güç"
	// Hafızadaki string (byte dizisi): [71, 195, 188, 195, 167]
	//                                  G   ---ü---   ---ç---

	// 1. DÖNÜŞÜM: string -> []rune
	runeSlice := []rune(s)

	// Go, [195, 188]'i okudu ve rune 252 ('ü') yaptı.
	// Go, [195, 167]'yi okudu ve rune 231 ('ç') yaptı.

	fmt.Println(runeSlice)
	// Çıktı (sayı olarak): [71 252 231]

	fmt.Printf("Karakterler: %c\n", runeSlice)
	// Çıktı (karakter olarak): [G ü ç]
}
