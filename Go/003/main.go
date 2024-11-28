package main

import (
	"fmt"
	"math"
)

const (
	// to = 13195
	to = 600851475143
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
	n := genPrimeNumbers(int(math.Pow(to, .5)))

	var answer int

	for _, num := range n {
		if to%num == 0 {
			answer = num
		}
	}

	fmt.Println(answer)
}
