package main

import (
	"fmt"
	"math"
)

const (
	to = 100
)

func main() {
	var x, y, answer int

	for i := 1; i <= to; i++ {
		x += int(math.Pow(float64(i), 2))
	}

	for j := 1; j <= to; j++ {
		y += j
	}

	y = int(math.Pow(float64(y), 2))
	answer = y - x

	fmt.Println(answer)
}
