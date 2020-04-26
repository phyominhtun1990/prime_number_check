package main

import "fmt"

func prime(number int) {
	Prime := true
	if number == 0 || number == 1 {
		fmt.Printf(" %d is not a  prime number\n", number)
	} else {
		for i := 2; i <= number/2; i++ { // Sieve of Eratosthenes Algorithm
			if number%i == 0 { // (%)_Remainder
				fmt.Printf(" %d is not a  prime number\n", number)
				Prime = false
				break
			}
		}
		if Prime == true {
			fmt.Printf(" %d is a prime number\n", number)
		}
	}
}

func main() {
	var number int
	fmt.Print("Check your number is prime or not: ")
	fmt.Scan(&number)
	prime(number)
}
