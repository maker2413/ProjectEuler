package main

import (
	"fmt"
	"math/big"
)

const (
	gridSize = 20
)

// Taken from: https://stackoverflow.com/a/54952509
func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}

func main() {
	// If you think it through no matter what combination of right and down moves you
	// make you can only reach the bottom right corner with n right moves and n down
	// moves where n is the size of the grid.
	x := big.NewInt(2 * gridSize)

	// From there we can find the total amount of unique combinations of moves as:
	// (2n)!/(n!)^2
	x = factorial(x)
	y := factorial(big.NewInt(gridSize))

	// https://pkg.go.dev/math/big#Int.Exp
	y = y.Exp(y, big.NewInt(2), big.NewInt(0))

	fmt.Println(x.Div(x, y))
}
