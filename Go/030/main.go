package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	from = 32     // 2^5
	to   = 354294 // 9^5 * 6
)

func digitFifthPowers(i int) int {
	var sum int

	num := strconv.Itoa(i)

	for i := 0; i < len(num); i++ {
		c, _ := strconv.Atoi(string(num[i]))
		sum += int(math.Pow(float64(c), 5))
	}

	return sum
}

func main() {
	var answer int

	for i := from; i < to; i++ {
		if i == digitFifthPowers(i) {
			answer += i
		}
	}

	fmt.Println(answer)
}
