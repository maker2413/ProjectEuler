package main

import (
	"fmt"
	"math"
	"time"
)

const (
	limit = 600851475143
)

func solve(to int) int {
	answer := 1

	if to%2 == 0 {
		answer = 2
		to /= 2

		for to%2 == 0 {
			to /= 2
		}
	}

	i := 3
	for i <= int(math.Pow(float64(to), .5)) {
		if to%i == 0 {
			answer = i

			for to%i == 0 {
				to /= i
			}
		}

		i += 2
	}

	if to > answer {
		return to
	}

	return answer
}

func main() {
	start := time.Now()
	answer := solve(limit)
	elapsed := time.Since(start)

	fmt.Println(answer)
	fmt.Println("--------------------------------------------------")
	fmt.Printf("  %.6f seconds\n", elapsed.Seconds())
}
