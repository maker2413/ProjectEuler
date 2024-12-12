package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Pow(float64(n), .5)); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var answer int
	primes := []int{2}

	for p := 3; len(primes) < 1000; p += 2 {
		if isPrime(p) {
			primes = append(primes, p)
		}
	}

	var consecutivePrimes int

	for i := 2; consecutivePrimes <= 2; i++ {
		consecutivePrimes++
	}

	fmt.Println(answer)
}
