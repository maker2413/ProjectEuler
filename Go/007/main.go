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

func genPrimeNumbers(n int) int {
	var nthPrime int

	for i := 3; n > 1; i += 2 {
		if isPrime(i) {
			n--
			nthPrime = i
		}
	}

	return nthPrime
}

func main() {
	fmt.Println(genPrimeNumbers(10001))
}
