package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)
	if isPrime(number) {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not Prime")
	}
}
func isPrime(number int) bool {
	if number < 2 {
		return false
	}

	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
