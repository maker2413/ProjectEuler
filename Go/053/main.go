package main

import (
	"fmt"
	"math/big"
	"strconv"
)

const (
	to  = 100
	mil = 1000000
)

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}

	return n.MulRange(1, x.Int64())
}

func combinatoric(n, r *big.Int) *big.Int {
	numerator := factorial(n)

	denominator := new(big.Int).Mul(factorial(r), factorial(big.NewInt(n.Int64()-r.Int64())))

	return new(big.Int).Div(numerator, denominator)
}

func main() {
	var answer int

	for n := 1; n <= to; n++ {
		for r := 0; r <= n; r++ {
			num := combinatoric(big.NewInt(int64(n)), big.NewInt(int64(r)))
			if len(num.String()) > len(strconv.Itoa(mil)) || num.Int64() > mil {
				answer++
			}
		}
	}

	fmt.Println(answer)
}
