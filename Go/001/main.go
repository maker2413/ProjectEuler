package main

import (
	"fmt"
	"time"
)

const (
	limit = 1_000_000
)

func solve(to int) int {
	var answer int

	for i := 3; i < to; i++ {
		if i%3 == 0 || i%5 == 0 {
			answer += i
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
