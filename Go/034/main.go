package main

import (
	"fmt"
	"strconv"
)

func factorialSum(n int, factorials []int) int {
	str := strconv.Itoa(n)
	var sum int

	for s := 0; s < len(str); s++ {
		i, _ := strconv.Atoi(string(str[s]))

		sum += factorials[i]
	}

	return sum
}

func main() {
	var answer int

	// 0! equals 1: https://www.thoughtco.com/why-does-zero-factorial-equal-one-3126598
	factorialDigits := []int{1}

	for i := 1; i <= 9; i++ {
		factorialDigits = append(factorialDigits, factorialDigits[i-1]*i)
	}

	// 9!*7 == 2540160 (7 digits), which means 2540160 is our upper limit
	for i := 10; i < 1499999; i++ {
		if i == factorialSum(i, factorialDigits) {
			answer += i
		}
	}

	fmt.Println(answer)
}
