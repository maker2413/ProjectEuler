package main

import (
	"fmt"
	"math"
)

func divisorCount(n int) int {
	var total int

	for i := 1; i < int(math.Pow(float64(n), .5))+1; i++ {
		if n%i == 0 {
			total += 2
		}
	}

	return total
}

func main() {
	var answer int = 1
	var i int = 2

	for divisorCount(answer) < 500 {
		answer += i
		i++
	}

	fmt.Println(answer)
}
