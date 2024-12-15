package main

import (
	"fmt"
	"math"
)

const (
	to = 1000000
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
	var maxConsec int = -1
	primes := []int{2}

	for p := 3; primes[len(primes)-1] < to; p += 2 {
		if isPrime(p) {
			primes = append(primes, p)
		}
	}

	for i := 0; i < len(primes); i++ {
		var consec int
		for j := i; j < len(primes); j++ {
			consec += primes[j]
			if consec > to {
				break
			}

			if isPrime(consec) && consec > answer && j-i > maxConsec {
				maxConsec = j - i
				answer = consec
			}
		}
	}

	fmt.Println(answer)
}
