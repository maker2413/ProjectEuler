package main

import (
	"fmt"
	"math"
)

const (
	consec = 4
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

func primeFactorCount(n int, primes []int) int {
	factors := make(map[int]int)

	for n != 1 {
		for _, prime := range primes {
			if n%prime == 0 {
				factors[prime]++
				n /= prime
				break
			}
		}
	}

	return len(factors)
}

func main() {
	var answer int
	primes := []int{2}

	for p := 3; len(primes) < 100000; p += 2 {
		if isPrime(p) {
			primes = append(primes, p)
		}
	}

	var consecutivePrimes int

	for i := 2; consecutivePrimes < consec; i++ {
		if primeFactorCount(i, primes) == consec {
			consecutivePrimes++
			answer = i
		} else {
			consecutivePrimes = 0
		}
	}

	fmt.Println(answer - consec + 1)
}
