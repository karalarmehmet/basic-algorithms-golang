package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("enter a number")
	var num int
	fmt.Scan(&num)

	fmt.Println("enter a primes list that will be used to factor the number you input")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // "1 2 3 4 5\n"
	input = strings.TrimSpace(input)    //"1 2 3 4 5"
	str := strings.Split(input, " ")    //["1" "2" "3" "4" "5"]

	a := []int{}
	for _, val := range str {
		b, _ := strconv.Atoi(val)
		a = append(a, b)
	}
	if primeNum(a) {
		res := factor(num, a)
		fmt.Println(res)

	} else {
		fmt.Println("there is a non-prime in the list")
	}

}

func factor(n int, a []int) []int {
	res := []int{}
	for _, val := range a {
		for n%val == 0 {
			res = append(res, val)
			n = n / val
		}
	}
	if n > 1 {
		res = append(res, n)
	}
	return res
}

func primeNum(p []int) bool {
	for _, val := range p {
		if val < 2 {
			return false
		} else if val > 2 && val%2 == 0 {
			return false
		}
		for i := 3; i*i <= val; i += 2 {
			if val%i == 0 {
				return false
			}
		}
	}
	return true
}
