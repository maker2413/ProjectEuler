package main

import (
	"fmt"
	"math"
	"strconv"
)

func removeLeftChar(s string) string {
	return s[1:]
}

func removeRightChar(s string) string {
	return s[:len(s)-1]
}

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

func isTruncatablePrime(n int) bool {
	if !isPrime(n) {
		return false
	}

	for str := strconv.Itoa(n); len(str) > 0; str = removeLeftChar(str) {
		i, _ := strconv.Atoi(str)

		if !isPrime(i) {
			return false
		}
	}

	for str := strconv.Itoa(n); len(str) > 0; str = removeRightChar(str) {
		i, _ := strconv.Atoi(str)

		if !isPrime(i) {
			return false
		}
	}

	return true
}

func sumTruncatablePrimes() int {
	var truncatablePrimes, sum int

	for i := 10; truncatablePrimes < 11; i++ {
		if isTruncatablePrime(i) {
			truncatablePrimes++
			sum += i
		}
	}

	return sum
}

func main() {
	answer := sumTruncatablePrimes()

	fmt.Println(answer)
}
