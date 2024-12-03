package main

import (
	"fmt"
	"math"
)

const (
	digitCount = 1000
)

func nthFib(element int) int {
	phiN := math.Pow(math.Phi, float64(element))
	phiNegativeN := -math.Pow(math.Phi, -float64(element))

	num := (phiN - phiNegativeN) / math.Sqrt(5)

	return int(num)
}

func nDigitNumberFib(n int) float64 {
	return (float64(n-1)*math.Log(10) + math.Log(5)/2) / math.Log(math.Phi)
}

func main() {
	answer := math.Ceil(nDigitNumberFib(digitCount))

	fmt.Println(answer)
}
