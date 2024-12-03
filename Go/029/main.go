package main

import (
	"fmt"
	"math/big"
)

const (
	terms = 100
)

func main() {
	var powers []big.Int

	for a := 2; a <= terms; a++ {
		for b := 2; b <= terms; b++ {
			var result big.Int
			result.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
			powers = append(powers, result)
		}
	}

	answer := make(map[string]int)

	for _, p := range powers {
		answer[p.String()]++
	}

	fmt.Println(len(answer))
}
