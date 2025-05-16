package main

import (
	"fmt"
	"math/big"
	"strconv"
)

const (
	limit = 100
)

func SumDigits(num string) int {
	var answer int

	for i := range num {
		digit, _ := strconv.Atoi(string(num[i]))
		answer += digit
	}

	return answer
}

func main() {
	var answer int
	var z big.Int

	for x := range limit {
		for y := range limit {
			a := big.NewInt(int64(x))
			b := big.NewInt(int64(y))

			z.Exp(a, b, nil)
			sum := SumDigits(z.String())

			if sum > answer {
				answer = sum
			}
		}
	}

	fmt.Println(answer)
}
