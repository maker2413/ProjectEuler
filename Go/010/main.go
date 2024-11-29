package main

import (
	"fmt"
	"math"
)

const (
	to = 2000000
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

func genPrimeNumbers(n int) []int {
	var list []int

	list = append(list, 2)

	for i := 3; i < n; i += 2 {
		if isPrime(i) {
			list = append(list, i)
		}
	}

	return list
}

func main() {
	var answer int

	primes := genPrimeNumbers(to)

	for i := 0; i < len(primes); i++ {
		answer += primes[i]
	}

	fmt.Println(answer)
}
