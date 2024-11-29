package main

import "fmt"

const (
	gridSize = 20
)

func countPaths(x, y int) int {
	var moveCount int
	if x == 0 && y == 0 {
		return 1
	}
	if x > 0 {
		moveCount = countPaths(x-1, y)
	}
	if y > 0 {
		moveCount += countPaths(x, y-1)
	}

	return moveCount
}

func main() {
	fmt.Println(countPaths(gridSize, gridSize))
}
