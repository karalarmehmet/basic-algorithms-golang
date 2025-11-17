package main

import (
	"fmt"
	"math/rand"
)

func main() {

	secretNumber := rand.Intn(10) + 1

	var guess, attempts int

	fmt.Println("I had a secret number try to guess that number")

	for {
		fmt.Println("please enter your guess:")
		fmt.Scan(&guess)
		attempts++

		if guess < secretNumber {
			fmt.Println("try to guess a greater number ")
		} else if guess > secretNumber {
			fmt.Println("try to guess a smaller number ")
		} else {
			fmt.Printf("correts congrats ! you got the secretNumber = %d it with %d attempts", secretNumber, attempts)
			break
		}

	}

}
