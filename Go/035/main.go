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

func main() {
	var answer int

	fmt.Println(answer)
}
