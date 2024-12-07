package main

import (
	"fmt"
	"math"
)

const (
	to = 1000
)

func findPerimeters(n int) [][]int {
	var perimeters [][]int

	for a := 1.0; a < math.Ceil(float64(n/2)); a++ {
		for b := 1.0; b < math.Ceil(float64(n/2)); b++ {
			c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))

			if a+b+c == float64(n) {
				perimeters = append(perimeters, []int{int(a), int(b), int(c)})
			}
		}
	}

	return perimeters
}

func main() {
	var answer, solutions int

	for i := 1; i <= to; i++ {
		x := len(findPerimeters(i))
		if x >= solutions {
			solutions = x
			answer = i
		}
	}

	fmt.Println(answer)
}
