package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n == 2 || n == 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func main() {
	diag := []int{1}
	sideLength := 1
	currentMax := 1
	primes := 0
	percentage := float32(primes) / float32(len(diag))

	for percentage >= 0.10 || percentage == 0 {
		sideLength += 2

		for range 4 {
			currentMax += sideLength - 1
			diag = append(diag, currentMax)

			if isPrime(currentMax) {
				primes++
			}
		}

		percentage = float32(primes) / float32(len(diag))
		//fmt.Printf("Current Percentage: %.2f, Current SideLength: %d\n", percentage, sideLength)
	}

	fmt.Println("Answer:", sideLength)
}
