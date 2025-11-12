package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func controlPrime(n int) {

	if n < 2 {
		fmt.Printf(" %d is less than 2 ,numbers less than 2 cannot be prime number", n)
		return
		//Burada return kullanmamızın sebebi, n eğer 2'den küçükse fonksiyonun ilerleyen kısımlarını çalıştırmasını engellemek.
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("%d is not a prime number", n)
			return // you need to put this return just to make sure no more control to exit the program safely

		}
	}
	fmt.Printf("%d is a prime number", n)

}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the number you want to find out whether it is prime or not. ")

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)

		if err != nil || num < 0 {
			fmt.Println("Please enter a valid number")
			continue
		}
		controlPrime(num)
		break

	}
}
