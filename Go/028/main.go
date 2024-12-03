package main

import (
	"fmt"
)

const (
	gridSize = 1001
)

func sprial(grid *[gridSize][gridSize]int) {
	x, y := int(gridSize/2), int(gridSize/2)
	count := 1
	radius := 1
	grid[y][x] = count

	for x != gridSize-1 && y != 0 {
		for i := 0; i < radius; i++ {
			count++
			x++
			grid[y][x] = count
		}
		for i := 0; i < radius; i++ {
			count++
			y++
			grid[y][x] = count
		}
		radius++
		for i := 0; i < radius; i++ {
			count++
			x--
			grid[y][x] = count
		}
		for i := 0; i < radius; i++ {
			count++
			y--
			grid[y][x] = count
		}
		radius++
	}
	for i := 0; i < radius-1; i++ {
		count++
		x++
		grid[y][x] = count
	}
}

func sumDiagonals(grid *[gridSize][gridSize]int) int {
	var sum int

	for i := 0; i < (gridSize-1)/2; i++ {
		sum += grid[i][(gridSize-1)-i]
		sum += grid[i][i]
		sum += grid[(gridSize-1)-i][i]
		sum += grid[(gridSize-1)-i][(gridSize-1)-i]
	}

	sum += grid[(gridSize-1)/2][(gridSize-1)/2]

	return sum
}

func main() {
	var answer int

	var grid [gridSize][gridSize]int

	sprial(&grid)

	answer = sumDiagonals(&grid)

	fmt.Println(answer)
}
