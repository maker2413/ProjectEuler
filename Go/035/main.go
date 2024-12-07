package main

import (
	"fmt"
	"math"
	"strconv"
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

func rotate(s string) string {
	tmp := s + s

	return tmp[1 : len(s)+1]
}

func isCircularPrime(n int) bool {
	if !isPrime(n) {
		return false
	}

	str := strconv.Itoa(n)

	for rotatedStr := rotate(str); rotatedStr != str; rotatedStr = rotate(rotatedStr) {
		i, _ := strconv.Atoi(rotatedStr)

		if !isPrime(i) {
			return false
		}
	}

	return true
}

func main() {
	var circularPrimes []int

	for i := 1; i < to; i++ {
		if isCircularPrime(i) {
			circularPrimes = append(circularPrimes, i)
		}
	}

	fmt.Println(len(circularPrimes))
}
