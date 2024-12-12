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

func isGoldbachConjecture(number int, primes []int) bool {
	for _, n := range primes {
		if math.Sqrt(float64(number-n)/2) == math.Round(math.Sqrt(float64(number-n)/2)) {
			return true
		}
	}

	return false
}

func main() {
	var answer int
	primes := []int{2}

	for num := 3; answer == 0; num += 2 {
		if isPrime(num) {
			primes = append(primes, num)
		} else {
			if !isGoldbachConjecture(num, primes) {
				answer = num
			}
		}
	}

	fmt.Println(answer)
}
