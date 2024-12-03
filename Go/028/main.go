package main

import (
	"fmt"
)

const (
	gridSize = 5
)

func main() {
	var answer int

	var grid [gridSize][gridSize]int

	grid[int(gridSize/2)][int(gridSize/2)] = 1

	fmt.Println(answer)
	fmt.Println(grid)
}
