package main

import (
	"fmt"
	"math"
)

const (
	to = 1000
)

func main() {
	var answer int

	for a := 1; a < to/3-2; a++ {
		for b := a + 1; b < to/2; b++ {
			c := math.Pow(math.Pow(float64(a), 2)+math.Pow(float64(b), 2), .5)
			if float64(a)+float64(b)+c == float64(to) {
				answer = a * b * int(c)
			}
		}
	}

	fmt.Println(answer)
}
