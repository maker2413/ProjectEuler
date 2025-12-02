package main

import (
	"fmt"
	"math"
)

const (
	to = 28123
)

var (
	abundantNums     [to + 1]bool
	abundantNumsList []int
)

func checkAbundant(n int) bool {
	sum := 1

	for i := 2; i < int(math.Pow(float64(n), .5))+1; i++ {
		if n%i == 0 {
			j := n / i
			sum += i
			if i != j {
				sum += j
			}
		}
	}

	return sum > n
}

func checkSumOfAbundant(n int) bool {
	for j := range abundantNumsList {
		if abundantNumsList[j] < n {
			if abundantNums[n-abundantNumsList[j]] {
				return true
			}
		}
	}

	return false
}

func main() {
	var answer int

	for i := 6; i <= to; i++ {
		abundantNums[i] = checkAbundant(i)
		if abundantNums[i] {
			abundantNumsList = append(abundantNumsList, i)
		}
	}

	for i := 1; i <= to; i++ {
		if !checkSumOfAbundant(i) {
			answer += i
		}
	}

	fmt.Println("Answer:", answer)
}
