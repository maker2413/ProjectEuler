package main

import "fmt"

const (
	to = 200
)

func countRecur(coins []int, n, total int) int {
	// if total is zero a solution has been found
	if total == 0 {
		return 1
	}

	// if total is below 0 or there is no more coins to try than return 0
	if total < 0 || n == 0 {
		return 0
	}

	return countRecur(coins, n, total-coins[n-1]) + countRecur(coins, n-1, total)
}

func possibleDenominations(coins []int, total int) int {
	// Not given any coins
	if total == 0 {
		return 0
	}

	return countRecur(coins, len(coins), total)
}

func main() {
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	//coins := []int{1, 2, 3}

	answer := possibleDenominations(coins, to)

	fmt.Println(answer)
}
