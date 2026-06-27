package main

import (
	"fmt"
	"time"
)

const (
	limit = 4_000_000
)

func solve(to int) int {
	answer := 0
	var n, m = 1, 1

	for m < to {
		m += n
		n = m - n
		if m%2 == 0 {
			answer += m
		}
	}

	return answer
}

func main() {
	start := time.Now()
	answer := solve(limit)
	elapsed := time.Since(start)

	fmt.Println(answer)
	fmt.Printf("%.6f seconds\n", elapsed.Seconds())
}
