package main

import (
	"fmt"
	"strconv"
	"time"
)

func palindromeCheck(n int) bool {
	num := strconv.Itoa(n)

	for i := range len(num) / 2 {
		if string(num[i]) != string(num[len(num)-1-i]) {
			return false
		}
	}

	return true
}

func solve() int {
	var answer int

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			n := i * j
			if palindromeCheck(n) && n > answer {
				answer = n
			}
		}
	}

	return answer
}

func main() {
	start := time.Now()
	answer := solve()
	elapsed := time.Since(start)

	fmt.Println(answer)
	fmt.Println("--------------------------------------------------")
	fmt.Printf("  %.6f seconds\n", elapsed.Seconds())
}
