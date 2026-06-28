package main

import (
	"fmt"
	"math"
	"time"
)

const (
	limit = 600851475143
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

func solve(to int) int {
	answer := 0

	n := genPrimeNumbers(int(math.Pow(float64(to), .5)))

	for _, num := range n {
		if to%num == 0 {
			answer = num
		}
	}

	return answer
}

func main() {
	start := time.Now()
	answer := solve(limit)
	elapsed := time.Since(start)

	fmt.Println(answer)
	fmt.Println("--------------------------------------------------")
	fmt.Printf("  %.6f seconds\n", elapsed.Seconds())
}
