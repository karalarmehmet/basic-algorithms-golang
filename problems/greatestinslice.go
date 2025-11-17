package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("enter a slice")
	reader := bufio.NewReader(os.Stdin) //create a reader
	slice, _ := reader.ReadString('\n') // "1 2 3 4\n"
	slice = strings.TrimSpace(slice)    // "1 2 3 4"
	subStr := strings.Split(slice, " ") //["1" "2" "3" "4" ]
	a := []int{}
	for i, _ := range subStr {
		res, _ := strconv.Atoi(subStr[i])
		a = append(a, res)
	}

	fmt.Println(greatestInSlice(a))

}

func greatestInSlice(a []int) int {
	if len(a) == 0 {
		return -1
	}
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}
