package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("enter a target")
	reader1 := bufio.NewReader(os.Stdin)
	target, _ := reader1.ReadString('\n')
	target = strings.TrimSpace(target)
	number, err := strconv.Atoi(target)
	if err != nil {
		fmt.Println("Invalid target")
	}

	a := []int{}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the slice leave spaces between numbers (ex: 2 7 11 15): ")
	slice, _ := reader.ReadString('\n')  //"2 7 11 15\n"
	slice = strings.TrimSpace(slice)     // delete the space characters like (\n \t)  "2 7 11 15\n" >> "2 7 11 15"
	strPart := strings.Split(slice, " ") //"2 7 11 15" → ["2", "7", "11", "15"]

	for _, val := range strPart {
		if val == "" {
			continue
		}
		num, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("invalid input")
		}
		a = append(a, num)
	}

	num1, num2 := find_two_that_sum(number, a)
	fmt.Println(num1, num2)
}

func find_two_that_sum(a int, b []int) (int, int) {
	for i := 0; i < len(b)-1; i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i]+b[j] == a {
				return b[i], b[j]

			}
		}
	}
	return -1, -1
}

// var sum int = 23
// a := [5]int{5, 8, 3, 2, 12}
// found := false
// for i := 0; i < 3; i++ {
// 	for j := i + 1; j < len(a); j++ {
// 		if a[i]+a[j] == sum {
// 			fmt.Println(a[i], a[j])
// 			found = true
// 		}
// 	}
// }
// if !found {
// 	fmt.Println("not in the list")
// }
