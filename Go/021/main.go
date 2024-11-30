package main

import (
	"fmt"
	"math"
)

const (
	to = 10000
)

func amicable(n int) int {
	divisors := 1

	for i := 2; i < int(math.Pow(float64(n), .5)); i++ {
		if n%i == 0 {
			divisors += i + (n / i)
		}
	}

	return divisors
}

func main() {
	var answer int

	for i := 1; i < to; i++ {
		j := amicable(i)
		if j > i && i == amicable(j) {
			answer += i + j
		}
	}

	fmt.Println(answer)
}
