package main

import "fmt"

func main() {

	arr := []int{34, 56, 23, 67, 12, 89, 2, 45}

	// En küçük ve en büyük sayıyı bulalım
	min, max := findMinMax(arr)

	// Sonuçları yazdıralım
	fmt.Printf("smallets num in the array: %d\n", min)
	fmt.Printf("greatest number in the array: %d\n", max)

}

func findMinMax(a []int) (int, int) {

	min, max := a[0], a[0]

	for _, num := range a {
		if num > max {
			max = num
		} else if num < min {
			min = num
		}

	}
	return min, max
}
