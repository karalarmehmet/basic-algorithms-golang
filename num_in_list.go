package main

import "fmt"

func num_in_list(a []int, target int) bool {
	for _, i := range a {
		if i == target {
			return true
		}
	}
	return false
}

func main() {
	a := []int{1, 2, 4, 6, 8, 56, 43, 21}
	var b int
	fmt.Println("Please enter a target")
	fmt.Scan(&b)
	res := num_in_list(a, b)
	if res == true {
		fmt.Println("the target is in the list")
	} else {
		fmt.Println("num is not in the list")
	}

}
