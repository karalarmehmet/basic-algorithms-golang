package main

import "fmt"

func main() {
	fmt.Println("create a list to calculate the sum of the list ")
	a := []int{}
	fmt.Scanln(&a)
	res := sum_list_of_numbers(a)
	fmt.Println(res)

}

func sum_list_of_numbers(a []int) int {
	sum := 0
	for _, val := range a {
		sum = sum + val
	}
	return sum
}
