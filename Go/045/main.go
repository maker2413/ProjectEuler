package main

import "fmt"

const (
	to = 50000
)

func triangleNumber(n int) int {
	return (n * (n + 1)) / 2
}

func pentagonalNumber(n int) int {
	return (n * (3*n - 1)) / 2
}

func hexagonalNumber(n int) int {
	return (n * (2*n - 1))
}

func main() {
	var answer int
	var pentNums, hexNums []int

	// Every Hexagonal number is also a triangle number
	// H_n = T_2n-1
	for h := 1; h <= to; h++ {
		hexNums = append(hexNums, hexagonalNumber(h))
	}

	for p := 1; p <= to; p++ {
		pentNums = append(pentNums, pentagonalNumber(p))
	}

	for h := 143; h < len(hexNums); h++ {
		for p := 165; p < len(pentNums); p++ {
			if hexNums[h] == pentNums[p] {
				answer = triangleNumber((2 * h) + 1)
			}
		}
	}

	fmt.Println("Answer:", answer)
}
