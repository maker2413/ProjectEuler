package main

import (
	"fmt"
	"math"
)

const (
	digitCount = 1000
)

func nDigitNumberFib(n int) float64 {
	return (float64(n-1)*math.Log(10) + math.Log(5)/2) / math.Log(math.Phi)
}

func main() {
	answer := math.Ceil(nDigitNumberFib(digitCount))

	fmt.Println("Answer:", answer)
}
