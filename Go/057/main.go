package main

import (
	"fmt"
	"math/big"
)

const (
	limit = 1000
)

var answer int

func RecursiveFraction(i int) (*big.Int, *big.Int) {
	if i == 1 {
		return big.NewInt(7), big.NewInt(5)
	}
	if i <= 0 {
		return big.NewInt(3), big.NewInt(2)
	}

	newN := big.NewInt(1)
	newD := big.NewInt(1)

	n, d := RecursiveFraction(i - 1)
	oldN, oldD := RecursiveFraction(i - 2)

	newN.Mul(big.NewInt(2), n)
	newN.Add(newN, oldN)

	newD.Mul(big.NewInt(2), d)
	newD.Add(newD, oldD)

	return newN, newD
}

func main() {
	numerators := make([]*big.Int, limit)
	denominators := make([]*big.Int, limit)

	numerators[0] = big.NewInt(3)
	numerators[1] = big.NewInt(7)
	denominators[0] = big.NewInt(2)
	denominators[1] = big.NewInt(5)

	for i := 2; i < limit; i++ {
		n := big.NewInt(1)
		d := big.NewInt(1)

		n.Mul(big.NewInt(2), numerators[i-1])
		n.Add(n, numerators[i-2])
		d.Mul(big.NewInt(2), denominators[i-1])
		d.Add(d, denominators[i-2])

		numerators[i] = n
		denominators[i] = d

		if len(n.String()) > len(d.String()) {
			answer++
		}
	}

	fmt.Println(answer)
}
